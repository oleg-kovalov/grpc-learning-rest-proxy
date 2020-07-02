package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-learning-rest-proxy/src/github.com/oleg.kovalov/grpc-learning-rest-proxy/user/userpb"
	"log"
)

func main() {

	fmt.Println("[client] starting User client ..")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// closes the resource at the very end of the program
	defer cc.Close()

	c := userpb.NewUserServiceClient(cc)

	listUsersCall(c)

}

func listUsersCall(c userpb.UserServiceClient) {
	fmt.Printf("[client] calling ListUsers RPC\n")


	req := &userpb.ListUsersRequest{}
	res, err := c.ListUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("received error for ListUsers RPC call: %v\n", err)
	}

	fmt.Printf("Users List:\n")
	for _,user := range res.GetUsers() {
		fmt.Printf("\t%v\n",user)
	}

}







