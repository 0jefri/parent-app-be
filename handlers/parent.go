package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/parent-app-be/models"
	"github.com/parent-app-be/pkg/firebase"
)

func ParentDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil token dari header Authorization
	idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if idToken == "" {
		http.Error(w, "Missing ID token", http.StatusBadRequest)
		return
	}

	// Verifikasi token
	token, err := firebase.VerifyFirebaseToken(idToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil UID dari token
	firebaseUID := token.UID

	// Cari parent berdasarkan firebase_uid
	var parent models.Parent
	if err := models.DB.Where("firebase_uid = ?", firebaseUID).First(&parent).Error; err != nil {
		http.Error(w, "Parent not found", http.StatusNotFound)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Parent detail retrieved",
		"data":    parent,
	})
}
