package config

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/bmizerany/pat"
	"github.com/gorilla/websocket"
)

type server struct {
	Router   http.Handler
	Upgrader websocket.Upgrader
	Views    *jet.Set
}

func MakeServer() *server {
	var u = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	var v = jet.NewSet(
		jet.NewOSFileSystemLoader("./html"),
		jet.InDevelopmentMode(),
	)

	s := &server{
		Router:   pat.New(),
		Upgrader: u,
		Views:    v,
	}

	return s
}

func (s *server) RunServer() {
	mux := s.SetupRouter()

	log.Println("Staring web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
