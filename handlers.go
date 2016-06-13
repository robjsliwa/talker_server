package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pborman/uuid"
)

// AddRoom - data format for add room message
type AddRoom struct {
	User   string `json:"user"`
	Room   string `json:"room"`
	UserID string `json:"user_id"`
	RoomID string `json:"room_id"`
}

// ChatText - data format chat text
type ChatText struct {
	User string `json:"user"`
	Room string `json:"room"`
	Text string `json:"text"`
}

// create new room and possibly user if he does not exist
func addRoom(client *Client, data interface{}) {
	var room *Room
	var user *User
	var message Message
	var addRoomData AddRoom
	messageData := map[string]interface{}{}

	fmt.Println(data)

	mapstructure.Decode(data, &addRoomData)

	fmt.Printf("%#v\n", addRoomData)

	existingRoom, isPresent := mainStore.FindRoom(addRoomData.Room)
	if isPresent {
		room = existingRoom
	} else {
		room = NewRoom(addRoomData.Room)
		mainStore.AddRoom(*room)
		go room.run()
	}

	room.join <- client

	user = &User{
		ID:   uuid.New(),
		Name: addRoomData.User,
	}
	messageData["room"] = room
	messageData["user"] = user
	message.Name = "room add"
	message.Data = messageData
	room.messageForward <- message
}

func chatText(client *Client, data interface{}) {
	var room *Room
	var chatTextData ChatText
	var chatMessage Message

	fmt.Println(data)

	mapstructure.Decode(data, &chatTextData)

	fmt.Printf("%#v\n", chatTextData)

	existingRoom, isPresent := mainStore.FindRoom(chatTextData.Room)
	if isPresent {
		room = existingRoom
	} else {
		room = NewRoom(chatTextData.Room)
		mainStore.AddRoom(*room)
		go room.run()
	}

	chatMessage.Name = "chat message"
	chatMessage.Data = chatTextData

	room.messageForward <- chatMessage
}
