package models

type Question struct {
	QuestionID   string  `json:"question_id"`
	ExerciseID   string  `json:"exercise_id"`
	QuestionText string  `json:"question_text"`
	Points       float32 `json:"points"`
}
