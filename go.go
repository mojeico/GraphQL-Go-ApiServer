package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mojeico/GraphQL-Go-ApiServer/repository"
	"mojeico/GraphQL-Go-ApiServer/service"

	"github.com/graphql-go/graphql"
)

var (
	userRepo    = repository.NewUserRepository()
	userService = service.NewUserService(userRepo)
)

func main() {

	var documentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Document",
			Fields: graphql.Fields{
				"DocumentId": &graphql.Field{
					Type: graphql.Int,
				},
				"DocumentName": &graphql.Field{
					Type: graphql.String,
				},
				"DocumentNumber": &graphql.Field{
					Type: graphql.Int,
				},
				"DocumentExpDate": &graphql.Field{
					Type: graphql.String,
				},
				"UserId": &graphql.Field{
					Type: graphql.Int,
				},
			},
			Description: "User document",
		},
	)

	var userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:       "User",
			Interfaces: nil,
			Fields: graphql.Fields{
				"UserId": &graphql.Field{
					Type: graphql.Int,
				},
				"UserName": &graphql.Field{
					Type: graphql.String,
				},
				"UserPassword": &graphql.Field{
					Type: graphql.String,
				},
				"UserDocument": &graphql.Field{
					Type: documentType,
				},
			},
			IsTypeOf:    nil,
			Description: "",
		},
	)

	userFields := graphql.Fields{
		"getUserById": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: userService.GetUserById(),
		},

		"userList": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "Get User List",
			Resolve:     userService.GetAllUser(),
		},

		"deleteUser": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve:     userService.DeleteUser(),
			Description: "Delete User by id",
		},
	}

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        userType,
				Description: "Create a new User",
				Args: graphql.FieldConfigArgument{
					"UserId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"UserName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userService.CreateUser(),
			},
		},
	})

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: userFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	/*queryAllUsers := `
	    {
	        list {
	            UserId
	            UserName
				UserDocument{
					DocumentName
				}
	        }
	    }
	`*/
	/*
	   	queryGetUserById := `
	       {
	           getUserById(id:1) {
	               UserId
	               UserName
	   			UserDocument{
	   				DocumentName
	   			}
	           }
	       }
	   `
	*/

	queryCreate := `
    mutation {
        createUser(UserId: 5, UserName:"test") {
            UserId
			UserName
        }
    }
`

	params := graphql.Params{Schema: schema, RequestString: queryCreate}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}

}
