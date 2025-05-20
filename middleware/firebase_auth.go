package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/parent-app-be/models"
	"github.com/parent-app-be/pkg/firebase"
)

type contextKey string

const ContextKeyParentID contextKey = "parentID"

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		if idToken == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// ✅ Gunakan fungsi dari package firebase
		token, err := firebase.VerifyFirebaseToken(idToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		uid := token.UID

		var parent models.Parent
		if err := models.DB.Where("firebase_uid = ?", uid).First(&parent).Error; err != nil {
			http.Error(w, "Parent not found", http.StatusUnauthorized)
			return
		}

		// Inject parentID ke context
		ctx := context.WithValue(r.Context(), ContextKeyParentID, parent.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
