package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/database/models"
)

// swagger:route POST /api/exercises exercises createExercise
// Creates a new exercise.
// responses:
//   201: exerciseResponse

// swagger:parameters createExercise
type createExerciseParams struct {
	// in:body
	Body models.Exercise
}

// swagger:response exerciseResponse
type exerciseResponseWrapper struct {
	// in:body
	Body models.Exercise
}

func (a *APIHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var exercise models.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call a function to save the exercise to DynamoDB
	err = a.dbConn.CreateExercise(&exercise)
	if err != nil {
		log.Printf("Failed to create exercise %s", err)
		http.Error(w, "Failed to create exercise", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exercise)
}

func (a *APIHandler) GetAllExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call a function to get all exercises from DynamoDB
	exercises, err := a.dbConn.GetAllExercises()
	if err != nil {
		http.Error(w, "Failed to retrieve exercises", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}
