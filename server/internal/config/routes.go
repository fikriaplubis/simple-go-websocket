package config

import (
	"net/http"

	"github.com/bmizerany/pat"

	"simple-go-websocket/server/internal/services"
	"simple-go-websocket/server/internal/transports"
)

func (s *server) SetupRouter() http.Handler {
	mux := pat.New()

	service := services.NewService(&s.Upgrader, s.Views)
	handler := transports.NewHandler(service)

	mux.Get("/", http.HandlerFunc(handler.Home))
	mux.Get("/ws", http.HandlerFunc(handler.WsEndpoint))

	return mux
}
