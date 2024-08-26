package main

import (
	"log"
	"net/http"

	"openim/handlers"
	"openim/server"
)

func main() {
	go server.StartWebSocketServer()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/ws", handlers.WebSocketHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
