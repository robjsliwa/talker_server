package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pborman/uuid"
)

// AddRoom - data format for add room message
type AddRoom struct {
	User string `json:"user"`
	Room string `json:"room"`
}

// create new room and possibly user if he does not exist
func addRoom(client *Client, data interface{}) {
	var room Room
	var user User
	var message Message
	var addRoomData AddRoom
	messageData := map[string]interface{}{}

	fmt.Println(data)

	mapstructure.Decode(data, &addRoomData)

	fmt.Printf("%#v\n", addRoomData)

	room.ID = uuid.New()
	room.Name = addRoomData.Room
	user.ID = uuid.New()
	user.Name = addRoomData.User

	existingUser, isPresent := mainStore.FindUser(user)
	if isPresent {
		user = *existingUser
	}

	messageData["room"] = room
	messageData["user"] = user
	message.Name = "room add"
	message.Data = messageData
	client.send <- message
}
