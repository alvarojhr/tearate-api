package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoDBConnection) CreateTable(tableName string, attributeDefinitions []types.AttributeDefinition, keySchema []types.KeySchemaElement) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attributeDefinitions,
		KeySchema:            keySchema,
		BillingMode:          types.BillingModePayPerRequest,
		TableName:            aws.String(tableName),
	}

	_, err := db.Client.CreateTable(context.TODO(), input)
	return err
}

func (db *DynamoDBConnection) CreateTables() error {
	// Create the Exercises table
	err := db.CreateTable("Exercises", []types.AttributeDefinition{
		{
			AttributeName: aws.String("exercise_id"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}, []types.KeySchemaElement{
		{
			AttributeName: aws.String("exercise_id"),
			KeyType:       types.KeyTypeHash,
		},
	})
	if err != nil {
		return err
	}

	err = db.CreateTable("Students", []types.AttributeDefinition{
		{
			AttributeName: aws.String("student_id"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}, []types.KeySchemaElement{
		{
			AttributeName: aws.String("student_id"),
			KeyType:       types.KeyTypeHash,
		},
	})
	if err != nil {
		return err
	}

	// Create the Questions table
	err = db.CreateTable("Questions", []types.AttributeDefinition{
		{
			AttributeName: aws.String("question_id"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}, []types.KeySchemaElement{
		{
			AttributeName: aws.String("question_id"),
			KeyType:       types.KeyTypeHash,
		},
	})
	if err != nil {
		return err
	}

	// Create the Answers table
	err = db.CreateTable("Answers", []types.AttributeDefinition{
		{
			AttributeName: aws.String("answer_id"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}, []types.KeySchemaElement{
		{
			AttributeName: aws.String("answer_id"),
			KeyType:       types.KeyTypeHash,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
