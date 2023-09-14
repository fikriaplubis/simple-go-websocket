package app

import (
	"simple-go-websocket/server/internal/config"
)

func StartApplication() {
	server := config.MakeServer()
	server.RunServer()
}
