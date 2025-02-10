package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/juanMaAV92/user-auth-api/utils"
)

func TraceID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceID := c.Get(utils.HeaderTraceID)
		if traceID == "" {
			traceID = fiberUtils.UUIDv4()
			c.Set(utils.HeaderTraceID, traceID)
		}

		c.Locals("trace_id", traceID)

		c.Response().Header.Set(utils.HeaderTraceID, traceID)
		return c.Next()
	}
}
