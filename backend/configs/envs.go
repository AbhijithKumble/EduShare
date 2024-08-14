package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the configuration needed for the application
type Config struct {
	Port                      string
	Db                        string
	Host                      string
	JWT_SECRET                string
	JWT_EXPIRATION_IN_SECONDS int64
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

	if c.JWT_EXPIRATION_IN_SECONDS == 0 {
		return fmt.Errorf("environment variable 'JWT_EXPIRATION_IN_SECONDS' is 0")
	}

	return nil
}

// GetConfig initializes a new Config struct by loading environment variables
func GetConfig() (*Config, error) {

	err := godotenv.Load(".env")
  
	if err != nil {
		return nil, fmt.Errorf(" error loading .env file: %v", err)
	}

	expirationStr := os.Getenv("JWT_EXPIRATION_IN_SECONDS")

	expiration, err := strconv.ParseInt(expirationStr, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing JWT_EXPIRATION_IN_SECONDS: %v", err)
	}
	config := &Config{
		Port:                      os.Getenv("PORT"),
		Db:                        os.Getenv("DB"),
		Host:                      os.Getenv("HOST"),
		JWT_SECRET:                os.Getenv("JWT_SECRET"),
		JWT_EXPIRATION_IN_SECONDS: expiration, 
	}

	if err := config.checkEnvs(); err != nil {
		return nil, err
	}

	return config, nil
}
