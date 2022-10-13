package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"spider/proto/user"
	"time"
)

const (
	address = "127.0.0.1:12345"
)

func mainc() {
	fmt.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserList(ctx, &user.GetUserListRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
