package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketController struct {
	// socketService services.SocketService
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// func NewSocketController(service services.SocketService) *SocketController {
// 	return &SocketController{socketService: service}
// }
