// config.go defines and initializes the configuration struct.
// It is responsible for setting up the configuration for the application,
// such as loading environment variables or config files.
package config

import (
	"os"
	"strconv"
)

type Config struct {
	// ...
	AWSRegion    string
	CreateTables bool
}

func LoadConfig() (*Config, error) {
	// ...
	awsRegion := os.Getenv("AWS_REGION")
	createTables := os.Getenv("CREATE_TABLES")

	isCreatetable, err := strconv.ParseBool(createTables)

	return &Config{
		AWSRegion:    awsRegion,
		CreateTables: isCreatetable,
	}, err
}
