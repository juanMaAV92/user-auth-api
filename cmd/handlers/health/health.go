package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanMaAV92/user-auth-api/services/health"
	"net/http"
)

type Handler struct {
	service *health.Service
}

func NewHandler(service *health.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Check(ctx *fiber.Ctx) error {
	response := h.service.HealthCheck()
	return ctx.Status(http.StatusOK).JSON(response)
}
