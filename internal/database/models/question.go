package models

type Question struct {
	QuestionID   string  `json:"question_id" dynamodbav:"question_id"`
	ExerciseID   string  `json:"exercise_id" dynamodbav:"exercise_id"`
	QuestionText string  `json:"question_text" dynamodbav:"question_text"`
	Points       float32 `json:"points" dynamodbav:"points"`
}
