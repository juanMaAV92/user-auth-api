package cmd

import (
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/juanMaAV92/user-auth-api/cmd/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juanMaAV92/user-auth-api/config"
)

type Server struct {
	Fiber  *fiber.App
	Config *config.Config
	Logger *log.Logger
}

func NewServer(cfg config.Config) *Server {
	fiberServer := fiber.New(
		fiber.Config{
			AppName: cfg.AppName,
		})

	// Global Middleware
	fiberServer.Use(requestid.New(
		requestid.Config{
			Header: "X-Trace-ID",
			Generator: func() string {
				return utils.UUIDv4()
			},
			ContextKey: "trace_id",
		}))
	fiberServer.Use(middleware.CustomLogger)

	return &Server{
		Fiber:  fiberServer,
		Config: &cfg,
		Logger: log.Default(),
	}
}

func (s *Server) Start() <-chan error {

	s.Logger.Printf("Starting server on port %s", s.Config.Port)

	errChan := make(chan error, 1)

	go func() {
		if err := s.Fiber.Listen(":" + s.Config.Port); err != nil {
			errChan <- err
		}
	}()

	return errChan

}
