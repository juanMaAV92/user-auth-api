package routes

import (
	"github.com/juanMaAV92/user-auth-api/cmd"
	"github.com/juanMaAV92/user-auth-api/cmd/handlers"
)

// Handler for initialising the routes

func RegisterRoutes(app *cmd.Server) {

	app.Fiber.Get("/health-check", handlers.HealthCheck)

	api := app.Fiber.Group("/api")
	v1 := api.Group("/v1")

	// Auth
	authGroup := v1.Group("/auth")
	authGroup.Post("/login", handlers.HealthCheck)
	authGroup.Post("/logout", handlers.HealthCheck)
}
