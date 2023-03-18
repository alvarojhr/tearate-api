// dynamodb.go initializes the DynamoDB client and provides helper functions for connecting to DynamoDB.
package database

import (
	"context"
	"fmt"
	"log"

	"github.com/alvarojhr/tearate-api/internal/database/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (db *DynamoDBConnection) GetAllExercises() ([]models.Exercise, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Exercises"),
	}

	output, err := db.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var exercises []models.Exercise
	err = attributevalue.UnmarshalListOfMaps(output.Items, &exercises)
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (db *DynamoDBConnection) CreateQuestion(question *models.Question) error {
	// Generate a new unique ID for the question
	question.QuestionID = uuid.New().String()

	item, err := attributevalue.MarshalMap(question)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Questions"),
	}

	_, err = db.Client.PutItem(context.TODO(), input)
	return err
}

func (db *DynamoDBConnection) GetAllQuestions() ([]models.Question, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Questions"),
	}

	output, err := db.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var questions []models.Question
	err = attributevalue.UnmarshalListOfMaps(output.Items, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (db *DynamoDBConnection) GetQuestion(questionID string) (*models.Question, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Questions"),
		Key: map[string]types.AttributeValue{
			"question_id": &types.AttributeValueMemberS{Value: questionID},
		},
	}

	result, err := db.Client.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("question not found")
	}

	question := &models.Question{}
	err = attributevalue.UnmarshalMap(result.Item, question)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (db *DynamoDBConnection) CreateStudent(student models.Student) error {

	student.StudentID = uuid.New().String()

	item, err := attributevalue.MarshalMap(student)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Students"),
	}

	_, err = db.Client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return err
}

func (db *DynamoDBConnection) GetAllStudents() ([]models.Student, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Students"),
	}

	output, err := db.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var students []models.Student
	err = attributevalue.UnmarshalListOfMaps(output.Items, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (db *DynamoDBConnection) CreateAnswer(answer models.Answer) error {

	answer.AnswerID = uuid.New().String()

	item, err := attributevalue.MarshalMap(answer)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Answers"),
	}

	_, err = db.Client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return err
}

func (db *DynamoDBConnection) GetAllAnswers() ([]models.Answer, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Answers"),
	}

	output, err := db.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var answers []models.Answer
	err = attributevalue.UnmarshalListOfMaps(output.Items, &answers)
	if err != nil {
		return nil, err
	}

	return answers, nil
}
