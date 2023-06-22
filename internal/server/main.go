package server

import (
	conf "github.com/msadministrator/golang-grpc-example/internal/config"
	"github.com/msadministrator/golang-grpc-example/internal/pb"
	_logger "github.com/msadministrator/golang-grpc-example/internal/logger"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var logger = _logger.NewLogger("info", true)

type UnimplementedExampleEnrichmentServer struct {
}

type ExampleEnrichmentServer struct {
	pb.UnimplementedExampleEnrichmentServer
}

func newServer() *ExampleEnrichmentServer {
	s := &ExampleEnrichmentServer{}
	return s
}

func (e *ExampleEnrichmentServer) enrichIP(ip string) (*pb.ExampleEnrichedIndicator) {
	logger.Info(fmt.Sprintf("Enriching IP: %s", ip))
	return &pb.ExampleEnrichedIndicator{
		Indicator: ip,
		Type: "ip",
		Malicious: true,
	}
}

func (e *ExampleEnrichmentServer) ExampleEnrich(ctx context.Context, in *pb.ExampleEnrichmentRequest) (*pb.ExampleEnrichmentResponse, error) {
	// We need to parse the data from the request
	logger.Info("ExampleEnrichmentServer received a request to the ExampleEnrich method on gRPC server.")
	logger.Info("Request body:\n" + in.String())

	var objects []*pb.ExampleEnrichedIndicator
	for i := 0; i < len(in.GetIps()); i++ {
		logger.Info(fmt.Sprintf("Processing IP %s", in.GetIps()[i]))
		logger.Info(fmt.Sprintf("Processing object %d", i))
		// call some other method here
		o := e.enrichIP(in.GetIps()[i])
		objects = append(objects, o)
	}
	return &pb.ExampleEnrichmentResponse{
		Objects: objects,
	}, nil
}

func Run() {
	c, err := conf.New()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	// Create a TCP listener on configured listeningPort
	fmt.Printf("Listening on port %d\n", c.ListeningPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.ListeningPort))
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}
	// Create an array of gRPC options
	opts := []grpc.ServerOption{}
	// Create a new gRPC server with the options
	s := grpc.NewServer(opts...)
	// Register the server with the protobuf generated code
	pb.RegisterExampleEnrichmentServer(s, newServer())
	logger.Info(fmt.Sprintf("server listening at %v", lis.Addr()))
	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
