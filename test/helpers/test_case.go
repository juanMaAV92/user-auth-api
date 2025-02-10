package helpers

import (
	"github.com/juanMaAV92/user-auth-api/cmd"
	"github.com/labstack/echo/v4"
)

type HttpTestCase struct {
	TestName    string
	Request     Request
	RequestBody interface{}
	MockFunc    func(s *cmd.AppServer, c echo.Context)
	Expected    ExpectedResponse
}

type Request struct {
	Method string
	Url    string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}
