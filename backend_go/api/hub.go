package api

import "github.com/charmbracelet/log"

// Hub maintains the set of active clients
// and broadcasts messages to them.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Imbound messages from clients.
	broadcast chan []byte

	// Register requests from clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// Handle registering and unregistering clients.
		case client := <-h.register:
			log.Debug("Registering client!")
			h.clients[client] = true
		case client := <-h.unregister:
			log.Debug("Unregistering client!")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Debug("Unregistered client!")
			}

		// Handle broadcasting messages to clients.
		case message := <-h.broadcast:
			log.Debugf("Recieved message: %v\n", string(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
