package main

import (
	"context"
	"log"
	pb "service/internal/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendMessage() {
	messages, err := grpc.NewClient(*messageService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer messages.Close()
	c := pb.NewMessageServiceClient(messages)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.CreateMessage{Text: "Hello!", Number: "88005553535"})
	if err != nil {
		log.Printf("could not send message: %v", err)
		return
	}
	log.Printf("Responce: %v", r.GetOk())
}
