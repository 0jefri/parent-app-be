package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/parent-app-be/models"
	"github.com/parent-app-be/pkg/firebase"
	"gorm.io/gorm"
)

func FirebaseRegisterHandler(w http.ResponseWriter, r *http.Request) {
	idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if idToken == "" {
		http.Error(w, "Missing ID token", http.StatusBadRequest)
		return
	}

	token, err := firebase.VerifyFirebaseToken(idToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	firebaseUID := token.UID
	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)
	phone, _ := token.Claims["phone_number"].(string)

	var parent models.Parent
	err = models.DB.Where("firebase_uid = ?", firebaseUID).First(&parent).Error
	if err == nil {
		http.Error(w, "User already registered", http.StatusConflict)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	parent = models.Parent{
		FirebaseUID: firebaseUID,
		Email:       email,
		Name:        name,
		PhoneNumber: phone,
	}
	if err := models.DB.Create(&parent).Error; err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Register berhasil",
		"user":    parent,
	})
}

func FirebaseLoginHandler(w http.ResponseWriter, r *http.Request) {
	idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token, err := firebase.VerifyFirebaseToken(idToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	firebaseUID := token.UID
	var parent models.Parent
	err = models.DB.Where("firebase_uid = ?", firebaseUID).First(&parent).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login berhasil",
		"user":    parent,
	})
}
