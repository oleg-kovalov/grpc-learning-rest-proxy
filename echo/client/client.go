package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/oleg-kovalov/grpc-learning-rest-proxy/echo/echopb"
	"log"
)

func main() {

	fmt.Println("[client] starting Echo client ..")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// closes the resource at the very end of the program
	defer cc.Close()

	c := echopb.NewEchoServiceClient(cc)

	echoCall(c)

}


func echoCall(c echopb.EchoServiceClient) {
	fmt.Println("[client] making Echo call to server")

	req := &echopb.EchoRequest{
		Msg: "hello from client",
	}

	res, err := c.Echo(context.Background(), req)
	if err != nil {
		log.Fatalf("[client] error while calling Echo RPC : %v", err)
	}

	log.Printf("[client] Response from sever: %v", res.GetMsg())
}
