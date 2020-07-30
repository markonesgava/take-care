package config

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

type Configuration struct {
	Port   int
	Google GoogleConfig
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
}

func ProvideServices(container *dig.Container) error {
	return container.Provide(newConfiguration)
}

func newConfiguration() *Configuration {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("Port"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	configuration := &Configuration{
		Port: port,
		Google: GoogleConfig{
			ClientID:     os.Getenv("Google_ClientID"),
			ClientSecret: os.Getenv("Google_ClientSecret"),
		},
	}

	return configuration
}
