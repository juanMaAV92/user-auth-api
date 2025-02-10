package helpers

import (
	"github.com/juanMaAV92/user-auth-api/cmd"
)

func NewServer() *cmd.AppServer {
	testServer := cmd.NewServer()

	return testServer
}
