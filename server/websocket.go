package server

import (
	"log"
	"net/http"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	ActiveConnections[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			delete(ActiveConnections, conn)
			break
		}

		BroadcastMessage(msg)
	}
}
