# clean up
go mod tidy

# install
go get \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/google/uuid


# generate
#  --go_out=plugins=grpc   generate DTO and gRPC Server and Client stubs
#  --grpc-gateway_out      generate gRPC-gateway related stuff
#  --swagger_out           generate swagger documentation
protoc \
    -I proto -I . \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=. \
    --swagger_out=./openapi \
    proto/user/user.proto           # proto/echo/echo.proto


# run gRPC server
go run user/server/server.go

# run gRPC-gateway
go run gateway/gateway.go

# run gRPC client
go run user/client/client.go
