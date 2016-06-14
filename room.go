package main

import (
	"log"

	"github.com/pborman/uuid"
)

// Room - struct for rooms where users can meet
type Room struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	messageForward chan Message
	join           chan *Client
	leave          chan *Client
	clients        map[*Client]bool
}

func NewRoom(name string) *Room {
	return &Room{
		ID:             uuid.New(),
		Name:           name,
		messageForward: make(chan Message),
		join:           make(chan *Client),
		leave:          make(chan *Client),
		clients:        make(map[*Client]bool),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)

		case message := <-r.messageForward:
			for client := range r.clients {
				select {
				case client.send <- message:

				default:
					// some failure sending message
					delete(r.clients, client)
					close(client.send)
					log.Println("Disconnecting from client")
				}
			}
		}
	}
}
