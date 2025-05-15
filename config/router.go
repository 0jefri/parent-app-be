package config

import (
	"net/http"

	"github.com/parent-app-be/handlers"
)

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.FirebaseRegisterHandler)
	mux.HandleFunc("/login", handlers.FirebaseLoginHandler)

	return mux
}
