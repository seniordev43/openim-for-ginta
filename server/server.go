package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ActiveConnections = make(map[*websocket.Conn]bool)

func StartWebSocketServer() {
	log.Println("WebSocket server started...")
}

func BroadcastMessage(msg []byte) {
	for conn := range ActiveConnections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
			conn.Close()
			delete(ActiveConnections, conn)
		}
	}
}
