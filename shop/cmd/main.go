package main

import (
	"flag"
	"log"
	"time"
)

var (
	userService    = flag.String("user_service", "user:50051", "the address to connect to")
	messageService = flag.String("message_service", "messages:50052", "the address to connect to")
)

func main() {
	flag.Parse()
	log.Println(*userService, *messageService)
	for {
		sendMessage()
		id := createUser()
		if id == "" {
			log.Println("could not get name")
			time.Sleep(time.Second * 5)
			continue
		}
		log.Printf("created user with id: %s", id)
		name := getUser(id)
		if name == "" {
			log.Println("could not get name")
			time.Sleep(time.Second * 5)
			continue
		}
		log.Printf("got username: %s", name)
		time.Sleep(time.Second * 5)
	}
}
