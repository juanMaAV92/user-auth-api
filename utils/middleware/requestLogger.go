package middleware

import (
	"github.com/juanMaAV92/user-auth-api/utils"
	"github.com/juanMaAV92/user-auth-api/utils/log"
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

func RequestLogger(c *fiber.Ctx, logger log.Logger) error {

	start := time.Now()
	traceID := c.Locals("trace_id").(string)

	importantHeaders := []string{utils.HeaderAuthorization, utils.HeaderContentType, utils.HeaderUserAgent, utils.HeaderTraceID}
	filteredHeaders := make(map[string][]string)

	for _, key := range importantHeaders {
		if value := c.Get(key); value != "" {
			filteredHeaders[key] = []string{value}
		}
	}

	request := requestLog{
		Headers: filteredHeaders,
		Body:    string(c.Body()),
	}

	err := c.Next()
	duration := time.Since(start).String()

	response := responseLog{
		Status: c.Response().StatusCode(),
		Body:   string(c.Response().Body()),
	}

	logger.Info(c.Context(), traceID, "request", "",
		log.Fields(map[string]interface{}{
			"method":   c.Method(),
			"url":      c.OriginalURL(),
			"request":  request,
			"response": response,
			"duration": duration,
		}),
	)

	return err
}
