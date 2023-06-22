package main

import (
	conf "github.com/msadministrator/golang-grpc-example/internal/config"
	_logger "github.com/msadministrator/golang-grpc-example/internal/logger"
	"github.com/msadministrator/golang-grpc-example/internal/server"
	"flag"
)


func main() {
	var logger = _logger.NewLogger("info", true)
	config, err := conf.New()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	var logLevel string

	// These are command line switches
	flag.IntVar(
		&config.ListeningPort,
		"port",
		3334,
		"Port to listen on",
	)
	flag.StringVar(
		&logLevel,
		"log-level",
		"info",
		"Logging level",
	)

	flag.Parse()

	logger.Info("Starting " + conf.AppName + " version " + conf.Version)

	// TODO: Add a shutdown channel
	// config.shutdown = make(chan os.Signal, 1)
	// signal.Notify(config.shutdown, syscall.SIGINT, syscall.SIGTERM)

	server.Run()
	logger.Info("Stopping " + conf.AppName + " version " + conf.Version)
}
