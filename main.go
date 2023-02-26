// RESTful API
// creating, reading, updating, and deleting user profile

// $ go build
// $ pardon_v1.exe
// => http://localhost:8080/users 에서 확인 가능

// or do 
// $ go run main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var users = make(map[string]User)
var logger *logrus.Logger

func init() {
	// Create a new logger instance
	logger = logrus.New()

	// Set the logger output to stdout
	logger.SetOutput(os.Stdout)

	// Set the logger level to info
	logger.SetLevel(logrus.InfoLevel)

	// Set the logger format to JSON
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Creating user...")

	// Decode the request body into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.WithError(err).Error("Failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	if _, ok := users[user.ID]; ok {
		logger.Warn("User already exists")
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	// Add the user to the map
	users[user.ID] = user

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	logger.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User created successfully")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Retrieving user...")

	// Get the user ID from the request URL
	vars := mux.Vars(r)
	userID := vars["id"]

	// Check if the user exists
	user, ok := users[userID]
	if !ok {
		logger.WithFields(logrus.Fields{
			"user_id": userID,
		}).Warn("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Send the user profile data in the response body
	json.NewEncoder(w).Encode(user)

	logger.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User retrieved successfully")
}

func main() {
	// Create a new Gorilla mux router
	r := mux.NewRouter()

	// Register the CreateUser and GetUser handlers
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")

	// Set up the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logger.Info("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}


/*
 TODO
 1. Add error handling - The current code does not handle errors that may occur 
 during request processing, such as invalid input data or database errors. 
 You can add error handling code using if statements or panic statements 
 to gracefully handle errors and return appropriate  HTTP response codes and error messages.

 2. Use a database - The current code stores user profiles in a map variable in memory, 
 which is not suitable for production use. You can use a database such as MySQL, PostgreSQL, 
 or MongoDB to store user data persistently and improve performance.

 3. Add authentication - The current code does not include any authentication mechanisms 
 to verify user identities or restrict access to protected endpoints. You can use a library 
 such as jwt-go or oauth2 to add authentication and authorization functionality.

 4. Use a web framework - The current code uses the standard http package, which can be difficult 
 to use and requires a lot of boilerplate code. You can use a web framework such as Gin, Echo, or Beego 
 to simplify request handling, routing, middleware, and error handling.

 5. Add logging - The current code does not include any logging mechanisms to record server activity or errors. 
 You can use a logging library such as logrus or zap to log server events and errors to a file or a database.

 6. Use environment variables - The current code uses hard-coded values for the server address and port, 
 which can be difficult to manage in different environments (e.g. development, staging, production). 
 You can use environment variables to store configuration values and make the code more portable and flexible.

 7. Write unit tests - The current code does not include any unit tests to verify 
 the correctness of the implementation and prevent regressions. You can use a testing library 
 such as testing or goconvey to write unit tests for each endpoint and HTTP handler.

 */