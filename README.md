# golang-grpc-example

This repository contains a simple example of a gRPC server and client written in Go.

> Please note that I have just started to learn Golang so this is not a production ready example.

This project utilizes [Protocol Buffers (protobufs)](https://protobuf.dev/) over [gRPC](https://grpc.io/) to communicate between the server and client. The server is a simple gRPC server that implements the `Example` service defined in the `example.proto` file. The client is a simple command line application that sends a request to the server and prints the response.

If you are unfamiliar with gRPC the benefits of it are that it is a high performance, open-source universal RPC framework. It is based on HTTP/2 and uses protobufs to communicate between the server and client. It is also language agnostic, meaning that you can write a server in one language and a client in another.

The example client provided (see below) connects over gRPC to the remote server and runs a simple command on that server. The server then responds with a message that is returned to the client. All of these objects are defined in the `example.proto` file.

## Development

This project uses protobufs to communicate over gRPC. To generate the Go code from the protobufs, run the following command:

```bash
cd internal/pb
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative example.proto 
cd ..
cd ..
```

## Running the example

To run the server run the following command:

```bash
go run ./cmd/golang-grpc-example
```

Next, open up a new tab and we need to run the client. To do this, run the following command:

```bash
go run ./pkg/client
```
