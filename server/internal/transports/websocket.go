package transports

import (
	"log"
	"net/http"

	"simple-go-websocket/server/internal/models"
	"simple-go-websocket/server/internal/services"
)

type Handler struct {
	Service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.Service.RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := h.Service.UpgradeConnection(w, r)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to endpoint")

	var response models.WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
}
