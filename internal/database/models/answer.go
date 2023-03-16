package models

type Answer struct {
	AnswerID   string `json:"answer_id"`
	QuestionID string `json:"question_id"`
	StudentID  string `json:"student_id"`
	Response   string `json:"response"`
}
