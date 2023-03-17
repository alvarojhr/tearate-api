package models

type Exercise struct {
	ExerciseID string `json:"exercise_id" dynamodbav:"exercise_id"`
	Name       string `json:"name" dynamodbav:"name"`
}
