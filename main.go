package main

import (
	"context"
	"github.com/juanMaAV92/user-auth-api/cmd"
	"github.com/juanMaAV92/user-auth-api/utils/log"
)

func main() {
	server := cmd.NewServer()
	server.Start()
	errChannel := server.Start()
	if err := <-errChannel; err != nil {
		server.Logger.Error(context.Background(), "", "", "Error while running", log.Field("error", err.Error()))
	}
}
