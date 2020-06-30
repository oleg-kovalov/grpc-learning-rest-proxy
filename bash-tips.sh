# clean up
go mod tidy

# install
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/google/uuid


# generate
protoc user/userpb/user.proto  --go_out=plugins=grpc:.

# generate grpc-rest-gateway
protoc user/userpb/user.proto  --grpc-gateway_out=logtostderr=true:.

# generate swagger
protoc user/userpb/user.proto --swagger_out=logtostderr=true:.



# run server
go run user/server/server.go

# run grpc-gateway
go run gateway/gateway.go

# run grpc client
go run user/client/client.go
