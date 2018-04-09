Before starting gRPC on golang you must read about protobuf which is developed by google.

### What is protobuf?

Protocol buffers are a flexible, fast, effcient and automated mechanism for serializing structured data from any language.
for more https://developers.google.com/protocol-buffers/docs/overview

### What is gPRC?

Like many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. By default, gRPC uses protocol buffers as the Interface Definition Language (IDL) for describing both the service interface and the structure of the payload messages
for more https://grpc.io/docs/guides/concepts.html

### Installation guide 
1. grab a protobuf as per your operating system from https://github.com/google/protobuf/releases .
2. Make environment variable for go packages.
3. run go get -u github.com/golang/protobuf/protoc-gen-go to install golang support.
4. run go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway for making reverse proxy or gateway over http or https.
5. run go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger for making a documentation for your service.
6. make GOPATH bin folder globally executable.
7. write a .proto file
8. `protoc -I. --go_out=plugins=grpc:. path/to/your_service.proto` for generating a stub file.
9. `protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. path/to/your_service.proto`
  It will generate a reverse proxy .pb.gw.go file
10. `protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. path/to/your_service.proto`
It will generate a swagger file which is contains given proto file descriptions.
11. checkout https://github.com/shakirmengrani/golang_examples/tree/master/grpc_example .

For further details https://github.com/grpc-ecosystem/grpc-gateway .
