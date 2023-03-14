package authentication

import (
	"context"
	"log"
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/v4"
    // "firebase.google.com/go/v4/auth"
	"cloud.google.com/go/firestore"
    "github.com/dgrijalva/jwt-go"
    "google.golang.org/api/iterator"
    "google.golang.org/api/option"

)

type User struct {
    UID			string    	`firestore:"u_id"`
    Name		string    	`firestore:"full_name"`
    Email		string    	`firestore:"email"`
	Level		int		`firestore:"avatar_level"`
    CreatedAt	time.Time 	`firestore:"created_at"`
    language	int 	`firestore:"language_code"`
}


func verifyTokenAndGetUID(token string) (string, error) {
	// 토큰 인증 확인

	// Initialize a Firebase app with your Firebase project credentials.
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return "", err
	}

	// Initialize a Firebase Auth client.
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
		return "", err
	}

	// Verify the Firebase authentication token and extract the UID.
	tokenData, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalf("error verifying id token: %v\n", err)
		return "", err
	}

	user_id := tokenData.UID
	return user_id, nil
}
  
func getUserByUID(uid string) (*User, error) {
	// returns `nil, nil` if the user not found.

	// Create a Firestore client.
	ctx := context.Background()
	projectID := "seesay"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Query the users collection for the user with the given UID.
	iter := client.Collection("users").Where("u_id", "==", uid).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err // Other Firestore error
	}

	// Extract the user's information from the Firestore document.
	user := &User{}
	err = doc.DataTo(user)
	if err != nil {
		return nil, err // Error parsing Firestore document
	}

	return user, nil
}
  
func createUserFromToken(token string) (*User, error) {
    // Verify the Firebase authentication token.
    ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create Firebase app: %v", err)
    }

	
    client, err := app.Auth(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create Firebase Auth client: %v", err)
    }

    decodedToken, err := client.VerifyIDToken(ctx, token)
    if err != nil {
        return nil, err
    }

    // Extract the user's information from the token.
    u_id := decodedToken.UID
    name := decodedToken.Claims["full_name"].(string)
    email := decodedToken.Claims["email"].(string)
	level := decodedToken.Claims["avatar_level"].(int)
    CreatedAt := decodedToken.Claims["creaged_at"].(time.Time)
	language := decodedToken.Claims["language_code"].(int)

    // Create a new User object with the user's information.
    user := &User{
		UID:	u_id,
		Name:	name,
		Email:	email,
		Level:	level,
		CreatedAt:	CreatedAt,
		language: language}

    return user, nil
}
  
func generateJWT(user *User) (string, error) {
	// Set the token's expiration time to a reasonable value (e.g., 1 hour).
	expirationTime := time.Now().Add(1 * time.Hour) // token expiration time = 한시간

	// Create a new JWT token with the user's UID as the subject.
	claims := jwt.MapClaims{
		"uid": user.UID,
		"exp": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secure secret key that only your backend server knows.
	// "your-secret-key": our own secret key.
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyFirebaseToken(ctx context.Context, token string) (string, error) {
    // Initialize Firebase app
    app, err := firebase.NewApp(ctx, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create Firebase app: %v", err)
    }

    // Verify ID token
    client, err := app.Auth(ctx)
    if err != nil {
        return "", fmt.Errorf("failed to create Firebase Auth client: %v", err)
    }

    idToken, err := client.VerifyIDToken(ctx, token)
    if err != nil {
        return "", fmt.Errorf("failed to verify Firebase ID token: %v", err)
    }

    // Extract user ID from token claims
    userID, ok := idToken.Claims["user_id"].(string)
    if !ok {
        return "", fmt.Errorf("Firebase ID token does not contain user ID")
    }

    return userID, nil
}


func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Receive the Firebase authentication token sent by the Flutter app after the user successfully logs in with Google.
	idToken := r.Header.Get("Authorization")
	ctx := context.Background()

	// Verify the Firebase authentication token using the Firebase Admin SDK for Go.
	uid, err := verifyFirebaseToken(ctx, idToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}


	// Check if the user already exists in the backend database by searching for the user's UID.
	user, err := getUserByUID(uid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// If the user does not exist, create a new user in the database with the user's UID and other relevant information (e.g., name, email, profile picture).
	if user == nil {
		user, err = createUserFromToken(idToken)
		if err != nil {
			return
		}
		err = addUserToFirestore(user)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Generate a new JWT token for the user using a secure token generation library such as JWT-go.
	jwtToken, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the JWT token to the Flutter app.
	w.Write([]byte(jwtToken))
}

func addUserToFirestore(user *User) error {
    // Replace with your own Firebase project ID and service account credentials
    ctx := context.Background()
    sa := option.WithCredentialsFile("/path/to/serviceAccountKey.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        return err
    }

    // Get a Firestore client
    client, err := app.Firestore(ctx)
    if err != nil {
        return err
    }
    defer client.Close()

    // Create a new document with a randomly generated ID
    newUserRef := client.Collection("users").NewDoc()
    // Set the user data on the document
    _, err = newUserRef.Set(ctx, *user) // pass the dereferenced user struct
    if err != nil {
        return err
    }

    return nil
}
  