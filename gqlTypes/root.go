package gqlTypes

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// Define a root query type
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Define fields for the root query
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get a user by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Implement resolver function to fetch user by ID
					// You can fetch the user from the database or any other data source
					// For simplicity, let's assume you have a GetUserByID function
					id, ok := params.Args["id"].(string)
					if !ok {
						return nil, nil
					}
					// Call a function to get the user from the database using the ID
					user, err := GetUserByID(id)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
		},
	},
)

// Define a mutation type
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        userInputType, // Assuming userType is already defined
			Description: "Create a new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Extract arguments from ResolveParams
				name, _ := params.Args["name"].(string)
				email, _ := params.Args["email"].(string)
				password, _ := params.Args["password"].(string)

				// Create a new user
				newUser := User{
					ID:       uuid.New(),
					Name:     name,
					Email:    email,
					Password: password,
				}

				// Return the created user
				return newUser, nil
			},
		},
	},
})

// Example function to fetch user from the database
func GetUserByID(id string) (interface{}, error) {
	// Implement your logic to fetch the user from the database
	// For this example, we'll just return a dummy user
	return User{
		ID:       uuid.UUID{},
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}, nil
}
