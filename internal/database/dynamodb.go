// dynamodb.go initializes the DynamoDB client and provides helper functions for connecting to DynamoDB.
package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBConnection struct {
	Client *dynamodb.Client
}

func NewDynamoDBConnection(region string) *DynamoDBConnection {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return &DynamoDBConnection{
		Client: svc,
	}
}
