# golang-grpc-example

This repository contains a simple example of a gRPC server and client written in Go.

> Please note that I have just started to learn Golang so this is not a production ready example.

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
