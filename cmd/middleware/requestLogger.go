package middleware

import (
	"github.com/juanMaAV92/user-auth-api/utils/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LogEntry struct {
	Method string `json:"method"`
	URL    string `json:"url"`

	Duration string `json:"duration"`
	Response string `json:"response"`
	TraceID  string `json:"trace_id"`
}

type requestLog struct {
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type responseLog struct {
	Body   string `json:"body"`
	Status int    `json:"status"`
}

func RequestLogger(c *fiber.Ctx, log *logger.Log) error {

	start := time.Now()
	traceID := c.Locals("trace_id").(string)

	request := requestLog{
		Headers: c.GetReqHeaders(),
		Body:    string(c.Body()),
	}

	err := c.Next()
	duration := time.Since(start).String()

	response := responseLog{
		Status: c.Response().StatusCode(),
		Body:   string(c.Response().Body()),
	}

	log.Info(traceID, "request",
		logger.Field("method", c.Method()),
		logger.Field("url", c.OriginalURL()),
		logger.Field("request", request),
		logger.Field("response", response),
		logger.Field("duration", duration),
	)

	return err
}
