package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func InitFirebaseApp() {
	keyPath := os.Getenv("FIREBASE_KEY_PATH")

	opt := option.WithCredentialsFile(keyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase init failed: %v", err)
	}

	AuthClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Auth init failed: %v", err)
	}
}

func VerifyFirebaseToken(idToken string) (*auth.Token, error) {
	return AuthClient.VerifyIDToken(context.Background(), idToken)
}
