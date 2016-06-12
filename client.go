package main

import "github.com/gorilla/websocket"

type FindHandler func(string) (Handler, bool)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
}

func (client *Client) Write() {
	for msg := range client.send {
		err := client.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
	client.socket.Close()
}

func (client *Client) Read() {
	var message Message
	for {
		err := client.socket.ReadJSON(&message)
		if err != nil {
			break
		}

		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

// NewClient - create new client
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}
