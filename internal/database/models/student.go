package models

type Student struct {
	StudentID    string `json:"student_id" dynamodbav:"student_id"`
	UniversityID string `json:"university_id" dynamodbav:"university_id"`
	Name         string `json:"name" dynamodbav:"name"`
}
