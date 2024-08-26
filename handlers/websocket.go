package handlers

import (
	"net/http"
	"openim/server"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	server.WebSocketHandler(w, r)
}
