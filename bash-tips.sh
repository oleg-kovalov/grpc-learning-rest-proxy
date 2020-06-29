# clean up
go mod tidy

# install
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go


# generate
protoc echo/echopb/echo.proto  --go_out=plugins=grpc:.

# generate grpc-rest-gateway
protoc echo/echopb/echo.proto --grpc-gateway_out=logtostderr=true:.

# generate swagger
protoc echo/echopb/echo.proto --swagger_out=logtostderr=true:.

