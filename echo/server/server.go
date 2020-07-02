package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/oleg.kovalov/grpc-learning-rest-proxy/echo/echopb"
	"log"
	"net"
)

type server struct{}

func (*server) Echo(ctx context.Context, request *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	message := request.GetMsg()
	fmt.Printf("[server] received message: %v\n", message)

	res := &echopb.EchoResponse{
		Msg: "Server responding with echo: " + message,
	}

	return res, nil
}

func main() {

	fmt.Printf("[server] starting Echo server .. \n")

	lis, err :=net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s:=grpc.NewServer()
	echopb.RegisterEchoServiceServer(s, &server{})

	if err:= s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}

}
