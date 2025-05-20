package config

import (
	"github.com/gorilla/mux"
	"github.com/parent-app-be/handlers"
	"github.com/parent-app-be/middleware"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Route publik (tidak pakai middleware)
	router.HandleFunc("/register", handlers.FirebaseRegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.FirebaseLoginHandler).Methods("POST")

	// 🔐 Route yang membutuhkan Firebase Auth
	auth := router.PathPrefix("/").Subrouter()
	auth.Use(middleware.FirebaseAuth)

	auth.HandleFunc("/parent/detail", handlers.ParentDetailHandler).Methods("GET")

	auth.HandleFunc("/children", handlers.GetChildren).Methods("GET")
	auth.HandleFunc("/children", handlers.CreateChild).Methods("POST")
	auth.HandleFunc("/children/{id}", handlers.UpdateChild).Methods("PUT")
	auth.HandleFunc("/children/{id}", handlers.DeleteChild).Methods("DELETE")

	return router
}
