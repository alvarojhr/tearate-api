package models

type Answer struct {
	AnswerID   string  `json:"answer_id" dynamodbav:"answer_id"`
	QuestionID string  `json:"question_id" dynamodbav:"question_id"`
	StudentID  string  `json:"student_id" dynamodbav:"student_id"`
	Response   string  `json:"response" dynamodbav:"response"`
	Rate       float32 `json:"points" dynamodbav:"points"`
}
