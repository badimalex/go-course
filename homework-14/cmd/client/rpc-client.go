package main

import (
	"fmt"
	"go-course/homework-14/pkg/messages"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	messagesToSend := []string{
		"Hello",
		"How are you?",
		"I'm using RPC!",
	}

	var reply bool

	for _, msg := range messagesToSend {
		err = client.Call("Server.Send", msg, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}

	var messages []messages.Message
	err = client.Call("Server.Messages", "", &messages)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		fmt.Printf("ID: %d, Time: %s, Content: %s\n", msg.ID, msg.Time, msg.Content)
	}

}
