package config

import (
	"os"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
	_logger "github.com/msadministrator/golang-grpc-example/internal/logger"
)

const AppName = "Golang gRPC Example"
const Version = "0.0.1"

var logger = _logger.NewLogger("info", true)

type ExamplePostgresConfig struct {
	credentials string
	host        string
	port        int
	name        string
	user        string
	password    string
}

type Configuration struct {
	ListeningPort 	int
	db            	ExamplePostgresConfig
	shutdown      	chan os.Signal
}

// NewConfig returns a new Config struct
func New() (*Configuration, error) {
	logger.Debug("Loading environment variables")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	config := &Configuration{}

	config.ListeningPort = config.GetEnvInt("LISTENING_PORT", 3334)
	// Add the rest of the environment variables here

	return config, nil
}

func (config *Configuration) GetEnvString(name string, fallback string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		logger.Warning(fmt.Sprintf("Environment variable %s not set, using default value %s", name, fallback))
		return fallback
	}
	return value
}

func (config *Configuration) GetEnvInt(name string, fallback int) int {
	value, _ := strconv.ParseInt(os.Getenv(name), 10, 64)
	if value == 0 {
		logger.Warning(fmt.Sprintf("Environment variable %s not set, using default value %d", name, fallback))
		return fallback
	}
	return int(value)
}

// printVersion displays the build version of the command
func (c *Configuration) printVersion() {
	logger.Info(AppName + " version " + Version)
}
