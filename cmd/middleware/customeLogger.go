package middleware

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LogEntry struct {
	Method   string              `json:"method"`
	URL      string              `json:"url"`
	Headers  map[string][]string `json:"headers"`
	Body     string              `json:"body"`
	Status   int                 `json:"status"`
	Duration string              `json:"duration"`
	Response string              `json:"response"`
	TraceID  string              `json:"trace_id"`
}

func CustomLogger(c *fiber.Ctx) error {

	start := time.Now()

	logEntry := LogEntry{
		Method:  c.Method(),
		URL:     c.OriginalURL(),
		Headers: c.GetReqHeaders(),
		Body:    string(c.Body()),
		TraceID: c.Locals("trace_id").(string),
	}

	err := c.Next()
	stop := time.Now()

	logEntry.Status = c.Response().StatusCode()
	logEntry.Response = string(c.Response().Body())
	logEntry.Duration = stop.Sub(start).String()

	logOutput, _ := json.Marshal(logEntry)

	log.Printf("%s\n", logOutput)

	return err
}
