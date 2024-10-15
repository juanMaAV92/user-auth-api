package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/juanMaAV92/user-auth-api/cmd/middleware"
	"github.com/juanMaAV92/user-auth-api/config"
	"github.com/juanMaAV92/user-auth-api/utils/logger"
)

type Server struct {
	Fiber  *fiber.App
	Config *config.Config
	Logger *logger.Log
}

func NewServer(cfg config.Config) *Server {
	fiberServer := fiber.New(
		fiber.Config{
			AppName: cfg.AppName,
		})

	// Logger
	appLogger := logger.NewLogger(cfg)

	// Global Middleware
	fiberServer.Use(requestid.New(
		requestid.Config{
			Header: "X-Trace-ID",
			Generator: func() string {
				return utils.UUIDv4()
			},
			ContextKey: "trace_id",
		}))
	fiberServer.Use(func(c *fiber.Ctx) error {
		return middleware.RequestLogger(c, appLogger)
	})

	return &Server{
		Fiber:  fiberServer,
		Config: &cfg,
		Logger: appLogger,
	}
}

func (s *Server) Start() <-chan error {

	s.Logger.Info("server", "Starting server on port "+s.Config.Port)

	errChan := make(chan error, 1)

	go func() {
		if err := s.Fiber.Listen(":" + s.Config.Port); err != nil {
			errChan <- err
		}
	}()

	return errChan

}
