package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the configuration needed for the application
type Config struct {
	Port       string
	Db         string
	Host       string
	JWT_SECRET string
}

var Envs = initConfig()

func initConfig() *Config {

	config, err := GetConfig()

	if err != nil {
		log.Fatal("Error in getting envs")
		return nil
	}
	return config
}

// checkEnvs verifies that all required environment variables are set
func (c *Config) checkEnvs() error {
	if c.Port == "" {
		return fmt.Errorf("environment variable 'Port' is empty")
	}
	if c.Db == "" {
		return fmt.Errorf("environment variable 'Db' is empty")
	}
	if c.Host == "" {
		return fmt.Errorf("environment variable 'Host' is empty")
	}
	if c.JWT_SECRET == "" {
		return fmt.Errorf("environment variable 'JWT_SECRET' is empty")
	}
	return nil
}

// GetConfig initializes a new Config struct by loading environment variables
func GetConfig() (*Config, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return nil, fmt.Errorf(" error loading .env file: %v", err)
	}

	config := &Config{
		Port:       os.Getenv("PORT"),
		Db:         os.Getenv("DB"),
		Host:       os.Getenv("HOST"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	if err := config.checkEnvs(); err != nil {
		return nil, err
	}

	return config, nil
}
