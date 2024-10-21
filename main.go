package main

import (
	"net/http"

	"github.com/charliekim2/multiplayer-typing-game/ws"
)

func main() {
	server := ws.NewServer()

	http.HandleFunc("/", server.Echo)
	http.ListenAndServe(":8080", nil)
}
