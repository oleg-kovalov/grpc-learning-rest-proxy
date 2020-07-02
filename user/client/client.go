package main

import (
	"context"
	"fmt"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc"
	"github.com/oleg.kovalov/grpc-learning-rest-proxy/user/userpb"
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

	//AddUser RPC unary call
	user, err := AddUser(c)

	//ListUsers RPC unary call
	ListUsers(c)

	//GetUser RPC unary call
	GetUser(c, user.GetId())
	//getUser(user.getId());

	//UpdateUser RPC unary call
	userToUpdate := &userpb.User{
		Id:        user.GetId(),
		FirstName: user.GetFirstName(),
		Email:     "jonny-d@yahoo.com",
		Longitude: user.GetLongitude(),
		Latitude:  user.GetLatitude(),
	}
	uMask := &field_mask.FieldMask{
		Paths: []string{"Email"},
	}
	UpdateUser(c, userToUpdate, uMask)

	//DeleteUser RPC unary call
	DeleteUser(c, user.GetId())

}

func UpdateUser(c userpb.UserServiceClient, userToUpdate *userpb.User, uMask *field_mask.FieldMask) (*userpb.User, error) {
	fmt.Printf("[client] calling UpdateUser RPC\n")

	req := &userpb.UpdateUserRequest{
		User:       userToUpdate,
		UpdateMask: uMask,
	}

	res, err := c.UpdateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] received error for UpdateUser RPC call: %v\n", err)
	}

	fmt.Printf("[client] updated user: %v", res.GetUser())
	return res.GetUser(), err
}

func DeleteUser(c userpb.UserServiceClient, id string) (*userpb.User, error) {
	fmt.Printf("[client] calling DeleteUser RPC\n")

	req := &userpb.DeleteUserRequest{UserId: id}

	res, err := c.DeleteUser(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] received error for DeleteUser RPC call: %v\n", err)
	}

	fmt.Printf("[client] deleted user: %v", res.GetUser())
	return res.GetUser(), err
}

func GetUser(c userpb.UserServiceClient, id string) (*userpb.User, error) {
	fmt.Printf("[client] calling GetUser RPC\n")

	req := &userpb.GetUserRequest{UserId: id}

	res, err := c.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] received error for GetUser RPC call: %v\n", err)
	}

	fmt.Printf("[client] found user by id: %v", res.GetUser())
	return res.GetUser(), err
}

func AddUser(c userpb.UserServiceClient) (*userpb.User, error) {
	fmt.Printf("[client] calling AddUser RPC\n")

	req := &userpb.AddUserRequest{
		User: &userpb.User{
			FirstName: "John",
			Email:     "john.doe@gmail.com",
			Latitude:  50.424858,
			Longitude: 30.506396,
		},
	}

	res, err := c.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] received error for AddUser RPC call: %v\n", err)
	}

	fmt.Printf("[client] added user: %v", res.GetUser())
	return res.GetUser(), err
}

func ListUsers(c userpb.UserServiceClient) ([]*userpb.User, error) {
	fmt.Printf("[client] calling ListUsers RPC\n")

	req := &userpb.ListUsersRequest{}
	res, err := c.ListUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] received error for ListUsers RPC call: %v\n", err)
	}

	fmt.Printf("Users List:\n")
	for _, user := range res.GetUsers() {
		fmt.Printf("\t%v\n", user)
	}

	return res.GetUsers(), err
}
