package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	echogw 	"github.com/oleg-kovalov/grpc-learning-rest-proxy/echo/echopb"
	usergw 	"github.com/oleg-kovalov/grpc-learning-rest-proxy/user/userpb"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50051", "gRPC server endpoint")
)

func run() error {
	fmt.Printf("[gateway] starting gRPC-REST reverse proxy ..\n")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := echogw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err2 := usergw.RegisterUserServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err2 != nil {
		return err2
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}