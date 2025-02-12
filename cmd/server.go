package cmd

import (
	"context"
	"github.com/gofiber/fiber/v2"
	config2 "github.com/juanMaAV92/user-auth-api/platform/config"
	"github.com/juanMaAV92/user-auth-api/services/health"
	"github.com/juanMaAV92/user-auth-api/utils/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	environment      = "ENVIRONMENT"
	localEnvironment = "local"
)

const (
	initServerStep = "init_server_step"
	shutdownStep   = "shutdown_server_step"
)

type AppServer struct {
	Fiber  *fiber.App
	Config *config2.Config
	Logger log.Logger
}

func NewServer() *AppServer {
	env := _getEnvironment()
	logger := log.New(config2.MicroserviceName, log.WithLevel(log.InfoLevel))
	cfg, err := config2.Load(env)
	if err != nil {
		logger.Error(context.Background(), "", initServerStep, "Error loading config", log.Field("error", err))
		panic("Error loading config: " + err.Error())
	}

	server := AppServer{
		Fiber: fiber.New(
			fiber.Config{
				AppName: cfg.MicroserviceName,
			}),
		Config: &cfg,
		Logger: logger,
	}
	services, err := server.initServices()
	if err != nil {
		server.Logger.Error(context.Background(), "", initServerStep, "Error initializing services", log.Field("error", err))
		panic("Error initializing services: " + err.Error())
	}
	configRoutes(&server, services)
	return &server
}

func (s *AppServer) initServices() (*AppServices, error) {
	healthService := health.NewService()

	return &AppServices{
		HealthService: healthService,
	}, nil
}

func (s *AppServer) Start() <-chan error {
	port := s.Config.HTTP.Port
	s.Logger.Info(context.Background(), "", initServerStep, "Starting server", log.Field("port", port))

	errC := make(chan error, 1)
	s.initGracefull(errC)

	go func() {
		if err := s.Fiber.Listen(":" + port); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC
}

func (s *AppServer) initGracefull(errChannel chan error) {
	gracefullTime := s.Config.HTTP.GracefullTime

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()
		s.Logger.Info(context.Background(), "", shutdownStep, "Shutdown signal received")
		_, cancel := context.WithTimeout(context.Background(), gracefullTime+(1*time.Second))
		defer func() {
			time.AfterFunc(gracefullTime, func() {
				if err := s.Fiber.Shutdown(); err != nil {
					errChannel <- err
				}
				s.Logger.Info(context.Background(), "", shutdownStep, "Shutdown completed")
				cancel()
				stop()
				close(errChannel)
			})
		}()

	}()
}

func _getEnvironment() string {
	env := localEnvironment
	if os.Getenv(environment) != "" {
		env = os.Getenv(environment)
	}
	return env
}
