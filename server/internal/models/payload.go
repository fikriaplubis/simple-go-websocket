package models

import (
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

// WsPayload defines the websocket request from the client
type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}
