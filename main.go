package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/parent-app-be/config"
	"github.com/parent-app-be/pkg/firebase"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init services
	firebase.InitFirebaseApp()
	config.InitDB()
	router := config.InitRouter()

	// Jalankan server
	port := os.Getenv("APP_PORT")
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
