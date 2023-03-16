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
	router.HandleFunc("/api/hello", apiHandler.HelloWorldHandler)

	apiHandler.Handler = router
	return apiHandler
}
