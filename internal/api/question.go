package api

import (
	"encoding/json"
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/database/models"
)

func (a *APIHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var question models.Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call a function to save the question to DynamoDB
	err = a.dbConn.CreateQuestion(&question)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func (a *APIHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call a function to get all questions from DynamoDB
	questions, err := a.dbConn.GetAllQuestions()
	if err != nil {
		http.Error(w, "Failed to retrieve questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
