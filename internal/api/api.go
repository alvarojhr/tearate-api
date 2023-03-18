package api

import (
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/database"
)

type APIHandler struct {
	dbConn *database.DynamoDBConnection
	http.Handler
}

func NewAPIHandler(dbConn *database.DynamoDBConnection) *APIHandler {
	apiHandler := &APIHandler{
		dbConn: dbConn,
	}

	router := http.NewServeMux()

	// Define your API routes here
	router.HandleFunc("/hello", apiHandler.HelloWorldHandler)
	router.HandleFunc("/exercises", apiHandler.HandleExercises)
	router.HandleFunc("/questions", apiHandler.HandleQuestions)
	router.HandleFunc("/students", apiHandler.HandleStudents)
	router.HandleFunc("/answers", apiHandler.HandleAnswers)

	apiHandler.Handler = router
	return apiHandler
}

func (a *APIHandler) HandleQuestions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.GetAllQuestions(w, r)
	case http.MethodPost:
		a.CreateQuestion(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *APIHandler) HandleExercises(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.GetAllExercises(w, r)
	case http.MethodPost:
		a.CreateExercise(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *APIHandler) HandleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.GetAllStudents(w, r)
	case http.MethodPost:
		a.CreateStudent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *APIHandler) HandleAnswers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.GetAllAnswers(w, r)
	case http.MethodPost:
		a.CreateAnswer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}