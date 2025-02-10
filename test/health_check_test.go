package tests

import (
	"fmt"
	"github.com/juanMaAV92/user-auth-api/config"
	"github.com/juanMaAV92/user-auth-api/test/helpers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	server := helpers.NewServer()
	request := helpers.Request{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/%s/health-check", config.MicroserviceName),
	}

	cases := []helpers.HttpTestCase{
		{
			TestName:    "HealthCheck - succeeded",
			Request:     request,
			RequestBody: nil,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   `{"status":"OK"}`,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			var requestBody string
			if test.RequestBody != nil {
				requestBody = test.RequestBody.(string)
			}

			req := httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, err := server.Fiber.Test(req, 5000)
			assert.NoError(t, err)
			assert.Equal(t, test.Expected.StatusCode, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(body), test.Expected.BodyPart)
		})
	}
}
