package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanMaAV92/user-auth-api/cmd/handlers/health"
	"github.com/juanMaAV92/user-auth-api/utils/middleware"
)

const (
	v1              = "/v1"
	healthCheckPath = "/health-check"
)

func configRoutes(server *AppServer, services *AppServices) {

	healthHandler := health.NewHandler(services.HealthService)
	configureMiddlewares(server)
	microserviceName := server.Config.MicroserviceName
	group := server.Fiber.Group(microserviceName)
	group.Get(healthCheckPath, healthHandler.Check)
}

func configureMiddlewares(server *AppServer) {
	server.Fiber.Use(middleware.TraceID())
	server.Fiber.Use(func(c *fiber.Ctx) error {
		return middleware.RequestLogger(c, server.Logger)
	})
}
