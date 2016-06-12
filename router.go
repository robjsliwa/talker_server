package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Handler - defines type for handler functions
type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Router - stores websocket handlers
type Router struct {
	rules map[string]Handler
}

// NewRouter - creates new Router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// Handle - adds handlers to the map
func (router *Router) Handle(msgName string, handler Handler) {
	router.rules[msgName] = handler
}

// FindHandler - looks up requested handler and returns it
func (router *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := router.rules[msgName]
	return handler, found
}

// ServeHTTP - upgrades http connection to socket and handles communication
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, router.FindHandler)
	go client.Write()
	client.Read()
}
