package main

import (
	"log"
	"net/http"

	"github.com/alvarojhr/tearate-api/internal/api"
	"github.com/alvarojhr/tearate-api/internal/config"
	"github.com/alvarojhr/tearate-api/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: No .env file found. Relying on system environment variables.")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}
	// Create a new DynamoDB connection
	dbConn := database.NewDynamoDBConnection(cfg.AWSRegion)

	if cfg.CreateTables {
		// Create the DynamoDB tables if they don't exist
		err = dbConn.CreateTables()
		if err != nil {
			log.Fatalf("Failed to create DynamoDB tables: %s", err.Error())
		}
	}

	apiHandler := api.NewAPIHandler(dbConn)

	addr := "localhost:8080"

	// Start the server
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, apiHandler))
}
