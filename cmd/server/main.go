package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// Cargar la configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	// Crear una nueva conexión a DynamoDB
	dbConn := database.NewDynamoDBConnection(cfg.AWSRegion)

	if cfg.CreateTables {
		// Crear las tablas de DynamoDB si no existen
		err = dbConn.CreateTables()
		if err != nil {
			log.Fatalf("Failed to create DynamoDB tables: %s", err.Error())
		}
	}

	apiHandler := api.NewAPIHandler(dbConn)

	// Obtener el puerto de la variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto si PORT no está definida
		log.Printf("Variable de entorno PORT no definida, usando el puerto por defecto %s", port)
	}

	// Dirección de escucha en 0.0.0.0 y el puerto especificado
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	// Iniciar el servidor
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, apiHandler))
}
