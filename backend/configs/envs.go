package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	DB     string
	CLIENT string
}

func checkEnvs(envString, name string) error {
	if envString == "" {
		return fmt.Errorf("environment variable '%s' is empty", name)
	}
	return nil
}

func getEnvs() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %v", err)
	}

	portString := os.Getenv("PORT")
	dbString := os.Getenv("DB")
	clientString := os.Getenv("CLIENT")

	if err := checkEnvs(portString, "PORT"); err != nil {
		return Config{}, err
	}

	if err := checkEnvs(dbString, "DB"); err != nil {
		return Config{}, err
	}

	if err := checkEnvs(clientString, "CLIENT"); err != nil {
		return Config{}, err
	}

	configs := Config{
		PORT:   portString,
		DB:     dbString,
		CLIENT: clientString,
	}

	return configs, nil
}

// Envs is the configuration loaded from environment variables
var Envs Config

func init() {
	var err error
	Envs, err = getEnvs()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
}

