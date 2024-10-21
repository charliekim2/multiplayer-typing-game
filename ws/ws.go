package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

type Server struct {
	clients map[*websocket.Conn]bool
}

func NewServer() *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
	}

	return &server
}

func (server *Server) Echo(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Bad connection")
		log.Panic(err)
	}

	server.clients[connection] = true // Save the connection using it as a key

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			log.Println("Error reading message: ", err)
			break
		}

		log.Println(string(message))
		server.WriteMessage(message)
	}

	delete(server.clients, connection) // Removing the connection

	err = connection.Close()
	if err != nil {
		log.Println("Could not close connection")
		log.Panic(err)
	}
}

func (server *Server) WriteMessage(message []byte) {
	for conn := range server.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Could not write message")
			log.Panic(err)
		}
	}
}
