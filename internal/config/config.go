// config.go defines and initializes the configuration struct.
// It is responsible for setting up the configuration for the application,
// such as loading environment variables or config files.
package config

import "os"

type Config struct {
	// ...
	AWSRegion string
}

func LoadConfig() (*Config, error) {
	// ...
	awsRegion := os.Getenv("AWS_REGION")

	return &Config{
		// ...
		AWSRegion: awsRegion,
	}, nil
}
