package main

import (
	"database/sql"
	// "database/sqlite"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./sql/users.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create users table if not exists
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
			password TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}

// User struct represents a user entity.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User struct
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	// fmt.println("test")
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", newUser.Name, newUser.Email)
	if err != nil {
		log.Println("Error inserting user:", err)
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Get the ID of the newly inserted user
	userID, _ := result.LastInsertId()

	// Set the ID in the response
	newUser.ID = int(userID)

	// Marshal the user as JSON
	responseJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request parameters
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	if !ok {
		http.Error(w, "User ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Convert the user ID to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Query the database to retrieve the user
	var user User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Error retrieving user:", err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// Marshal the user as JSON
	responseJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request parameters
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	if !ok {
		http.Error(w, "User ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Convert the user ID to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse the request body into a User struct
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Update the user in the database
	_, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", updatedUser.Name, updatedUser.Email, userID)
	if err != nil {
		log.Println("Error updating user:", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Set the ID in the response
	updatedUser.ID = userID

	// Marshal the updated user as JSON
	responseJSON, err := json.Marshal(updatedUser)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request parameters
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	if !ok {
		http.Error(w, "User ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Convert the user ID to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete the user from the database
	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		log.Println("Error deleting user:", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	responseJSON := []byte(`{"message":"User deleted successfully"}`)

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	// Initialize your database connection and routing here

	// Example route for creating a user

	// Start the server

	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

	port := 8080
	log.Printf("Server started on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
