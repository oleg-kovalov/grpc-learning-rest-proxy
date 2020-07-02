# grpc-learning-rest-proxy
This repository is an example of gRPC services interaction in Golang.\
This project contains a sample implementation of gRPC server and gRPC client \
as well as gRPC-gateway for interaction with generic REST clients.

It is also a companion to [grpc-learning-java](https://github.com/oleg-kovalov/grpc-learning-java) project written in Java.\
Projects could be used together to demonstrate language interoperability.

Example scenarious :
* Go gRPC server, Java gRPC client
* Java gRPC server, Go gRPC-gateway, generic REST client

## Installation
Use bash-tips.sh commands to clean, get dependencies, generate files from .proto and run components you need.

## Usage
Generate Go files from .proto file
```bash
$ protoc \
>    -I proto -I . \
>    --go_out=plugins=grpc:. \
>    --grpc-gateway_out=. \
>    --swagger_out=./openapi \
>    proto/echo/echo.proto  
```
Start Golang gRPC server
```bash
$ go run user/server/server.go
```
Start Golang gRPC client 
```bash
$ go run user/client/client.go
```
Start gRPC-gateway 
```bash
$ go run gateway/gateway.go
```

## Language interoperability
In many scenarious client and server could be differently implemented.\
gRPC being platform agnostic by design allows such communication, using same .proto files across all participants as the single source of truth.

In order to try such scenario current project could be used together with [grpc-learning-java](https://github.com/oleg-kovalov/grpc-learning-java) project written in Java.

Start Golang gRPC server
```bash
$ go run user/server/server.go
```

Start Java gRPC client
```java
user.client.Client.java
```

## gRPC-gateway
gRPC-gateway is a project aimed to support REST communication in a gRPC-centralized application.\
Thus gRPC could be added to existing RESTful application with backward compatibility for existing REST clients.\
This is achieved via additional reverse-proxy server which translates a RESTful HTTP API into gRPC.\
For more information please visit [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) repository.\
![gRPC-gateway architecture](https://docs.google.com/drawings/d/12hp4CPqrNPFhattL_cIoJptFvlAqm5wLQ0ggqI5mkCg/pub?w=749&amp;h=370)

to try scenario with gRPC-gateway current project could be used together with [grpc-learning-java](https://github.com/oleg-kovalov/grpc-learning-java).

Start Java gRPC server
```java
user.server.Server.java
```
Start Golang gRPC-gateway
```bash
$ go run user/gateway/gateway.go
```

You can make calls using a generic REST client.\
Each call would be intercepted by gRPC-gateway and translated to gRPC call that would be processed on Java gRPC server.
