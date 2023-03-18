package api

import (
	"encoding/json"
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/database/models"
)

func (a *APIHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call a function to create a new student in DynamoDB
	err = a.dbConn.CreateStudent(student)
	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (a *APIHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call a function to get all students from DynamoDB
	students, err := a.dbConn.GetAllStudents()
	if err != nil {
		http.Error(w, "Failed to retrieve students", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
