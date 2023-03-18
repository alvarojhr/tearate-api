package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/database/models"
)

func (a *APIHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var answer models.Answer
	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call a function to create a new answer in DynamoDB
	err = a.dbConn.CreateAnswer(answer)
	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	// Fetch the question details using the questionID from the answer
	question, err := a.dbConn.GetQuestion(answer.QuestionID)
	if err != nil {
		// Handle the error, e.g., send an error response
		return
	}

	rating, err := RateAnswer(question.QuestionText, question.Points, answer.Response)
	if err != nil {
		// Handle the error
	}

	log.Printf(rating)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func (a *APIHandler) GetAllAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call a function to get all answers from DynamoDB
	answers, err := a.dbConn.GetAllAnswers()
	if err != nil {
		http.Error(w, "Failed to retrieve answers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}
