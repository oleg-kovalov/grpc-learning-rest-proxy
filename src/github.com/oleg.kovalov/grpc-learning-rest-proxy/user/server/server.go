package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-learning-rest-proxy/src/github.com/oleg.kovalov/grpc-learning-rest-proxy/user/userpb"
	"log"
	"net"
)


type server struct{
	users  []*userpb.User

	//users := []*userpb.User {
	//	&userpb.User{
	//		Id: 1,
	//		FirstName: "John",
	//		Email: "john.doe@gmail.com",
	//		Latitude: 50.424858,
	//		Longitude: 30.506396,
	//	},
	//	&userpb.User{
	//		Id: 2,
	//		FirstName: "Vasya",
	//		Email: "vasyl.pupkin@gmail.com",
	//		Latitude: 64.181247,
	//		Longitude: -51.693769,
	//	},

	//}
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	fmt.Printf("[server] DeleteUser operation was invoked\n")

	var u *userpb.User
	tmp := s.users[:0]
	for _, user := range s.users {
		if user.GetId() != req.GetUserId() {
			tmp = append(tmp, user)
		} else {
			u = user
		}
	}
	s.users = tmp

	if u == nil {
		return nil, status.Errorf(codes.NotFound, "user with ID %q was not found", req.GetUserId())
	}

	res := &userpb.DeleteUserResponse{
		User: u,
	}
	return res, nil
}


func (s *server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	fmt.Printf("[server] UpdateUser operation was invoked\n")

	var u *userpb.User

	for _, user := range s.users {
		if user.GetId() == req.GetUser().GetId() {
			u = user
			break
		}
	}

	if u == nil {
		return nil, status.Errorf(codes.NotFound, "user with ID %q was not found", req.GetUser().GetId())
	}

	for _,path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "email":
			u.Email = req.GetUser().GetEmail()
		default :
			return nil, status.Errorf(codes.InvalidArgument, "could not update field %q on user", path)
		}
	}

	res := &userpb.UpdateUserResponse{
		User: u,
	}
	return res, nil
}


func (s *server) AddUser(ctx context.Context,req *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	fmt.Printf("[server] AddUser operation was invoked\n")

	userReq := req.GetUser()

	user := &userpb.User{
		Id: uuid.New().String(),
		FirstName: userReq.GetFirstName(),
		Email: userReq.GetEmail(),
		Latitude: userReq.GetLatitude(),
		Longitude: userReq.GetLongitude(),
	}

	s.users = append(s.users, user)

	res := &userpb.AddUserResponse{
		User: user,
	}
	return res, nil
}


func (s *server) GetUser(ctx context.Context,req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	fmt.Printf("[server] GetUser operation was invoked\n")

	for _, user := range s.users {
		if user.GetId() == req.GetUserId() {
			res := &userpb.GetUserResponse{
				User: user,
			}
			return res, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "user with ID %q was not found", req.GetUserId())
}


func (s *server) ListUsers(context.Context, *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	fmt.Printf("[server] ListUsers operation was invoked\n")

	res := &userpb.ListUsersResponse{
		Users: s.users,
	}
	return res, nil
}


func main() {

	fmt.Printf("[server] starting User server .. \n")

	lis, err :=net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s:=grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	if err:= s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v\n", err)
	}

}
