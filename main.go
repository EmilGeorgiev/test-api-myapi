package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = map[string]User{}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the user, like save it to a database
	users[user.Email] = user

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Create a new HTTP server
	mux := http.NewServeMux()

	// Register the handler function for the /users endpoint
	mux.HandleFunc("/users", createUserHandler)

	// Start the server
	http.ListenAndServe(":8080", mux)
}
