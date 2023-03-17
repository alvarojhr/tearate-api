// dynamodb.go initializes the DynamoDB client and provides helper functions for connecting to DynamoDB.
package database

import (
	"context"
	"log"

	"github.com/alvarojhr/tearate-api/internal/database/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
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

func (db *DynamoDBConnection) CreateExercise(exercise *models.Exercise) error {
	// Generate a new unique ID for the exercise
	exercise.ExerciseID = uuid.New().String()

	log.Print(exercise)

	item, err := attributevalue.MarshalMap(exercise)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Exercises"),
	}

	_, err = db.Client.PutItem(context.TODO(), input)
	return err
}
