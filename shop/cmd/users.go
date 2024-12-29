package main

import (
	"context"
	"log"
	pb "service/internal/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createUser() string {
	users, err := grpc.NewClient(*userService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return ""
	}
	defer users.Close()
	c := pb.NewUserServiceClient(users)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.CreateUserRequest{Name: "Muhammed Abdul", Email: "some-email-idk@gmail.com"})
	if err != nil {
		log.Printf("could not create user: %v", err)
		return ""
	}
	return r.GetUserId()
}

func getUser(id string) string {
	users, err := grpc.NewClient(*userService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return ""
	}
	defer users.Close()
	c := pb.NewUserServiceClient(users)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.GetUserRequest{UserId: id})
	if err != nil {
		log.Printf("could not get user: %v", err)
		return ""
	}
	return r.GetName()
}
