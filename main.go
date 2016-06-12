package main

import "net/http"

// Room - struct for rooms where users can meet
type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User - struct identifies valid user
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	// fake data store
	mainStore, _ = NewMockStore()

	router := NewRouter()

	router.Handle("room add", addRoom)
	router.Handle("chat message", chatText)

	http.Handle("/", router)
	http.ListenAndServe(":4652", nil)
}

/*

addRoom(userName, roomName)
addUserToRoom(userName, roomName) - if the user does not exist client will get
  new user name.  If client already remembers user it will pass that user name.
  Server needs to verify if the user exists or not and create a new user if needed.
  Server returns success or failure.  Client will forget user on failure?

*/
