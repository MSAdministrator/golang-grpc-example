package main

import (
	"github.com/msadministrator/golang-grpc-example/internal/pb"
	_logger "github.com/msadministrator/golang-grpc-example/internal/logger"
	conf "github.com/msadministrator/golang-grpc-example/internal/config"
	"context"
	"fmt"

	"google.golang.org/grpc"

	"flag"
	"os"
)


var logger = _logger.NewLogger("info", true)

func main() {
	var err error
	var logLevel string

	c, err := conf.New()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	flag.IntVar(
		&c.ListeningPort,
		"port",
		3334,
		"Port to send requests to",
	)
	flag.StringVar(
		&logLevel,
		"log-level",
		"info",
		"Logging level",
	)

	flag.Parse()

	logger.Info("Starting client")
	addr := fmt.Sprintf(":%v", c.ListeningPort)
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to server", err))
		os.Exit(1)
	}
	defer conn.Close()

	// Creating the data struct to send to the ExampleEnrichment server
	var ips []string
	ips = append(ips, "123.123.123.123")
	ips = append(ips, "8.8.8.8")

	// Creating the ExampleEnrichment client to send the data to
	// NewExampleEnrichmentClient is a generated function based on
	// the proto definition example.proto.
	client := pb.NewExampleEnrichmentClient(conn)

	// Creating a batch of IP addresses to enrich
	batch := &pb.ExampleEnrichmentRequest{
		Ips: ips,
	}

	// Calling the ExampleEnrich function on the server with our batch of Ips to enrich
	// ExampleEnrich is a function based on the `rpc ExampleEnrich (ExampleEnrichmentRequest) returns (ExampleEnrichmentResponse);`
	// definition in example.proto.
	resp, err := client.ExampleEnrich(
		context.Background(),
		batch,
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error calling ExampleEnrich", err))
	}
	logger.Info(fmt.Sprintf("Received a response ->", resp))
}
