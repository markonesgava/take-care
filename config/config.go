package config

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

type Configuration struct {
	Port    int
	Google  GoogleConfig
	MongoDB MongoDB
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
}

type MongoDB struct {
	ConnectionString string
}

func ProvideServices(container *dig.Container) error {
	return container.Provide(newConfiguration)
}

func newConfiguration() *Configuration {
	dir, _ := os.Getwd()
	log.Printf("base %s", dir)

	err := godotenv.Load("../.env") // fmt.Sprintf("../.env", dir))

	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	port, err := strconv.Atoi(os.Getenv("Port"))

	if err != nil {
		log.Fatalf("Error on get PORT %s", err)
	}

	configuration := &Configuration{
		Port: port,
		Google: GoogleConfig{
			ClientID:     os.Getenv("Google_ClientID"),
			ClientSecret: os.Getenv("Google_ClientSecret"),
		},
		MongoDB: MongoDB{
			ConnectionString: os.Getenv("MongoDB_ConnectionString"),
		},
	}

	return configuration
}
