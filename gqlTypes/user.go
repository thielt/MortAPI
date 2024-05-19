package mortapi

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// User struct represents a user entity.
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// Define GraphQL types for User
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.String},
			"name":     &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"password": &graphql.Field{Type: graphql.String},
		},
	},
)

// Define GraphQL input type for creating/updating User
var userInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
