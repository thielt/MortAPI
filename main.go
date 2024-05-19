package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thielt/mortapi/gqlTypes"
)

func main() {
	// Initialize the database connection
	InitDB()

	// Define your GraphQL schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    gqlTypes.RootQuery,    // Define your root query
		Mutation: gqlTypes.RootMutation, // Define your root mutation (if any)
	})
	if err != nil {
		log.Fatal("Failed to create schema:", err)
	}

	r := mux.NewRouter()

	// Handle GraphQL requests
	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		// Execute GraphQL query
		ExecuteQuery(w, r, schema)
	}).Methods("POST")

	port := 8080
	log.Printf("Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func ExecuteQuery(w http.ResponseWriter, r *http.Request, schema graphql.Schema) {
	// Parse GraphQL query from request body
	var query struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Execute the query against the schema
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query.Query,
		VariableValues: query.Variables,
	})

	// Encode the result as JSON and write it to the response
	json.NewEncoder(w).Encode(result)
}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./sql/users.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the database connection is valid by pinging it
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Create Users table if none found
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}
