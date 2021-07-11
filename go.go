package main

import (
	"github.com/graphql-go/handler"
	"log"
	"mojeico/GraphQL-Go-ApiServer/graphql/document"
	"mojeico/GraphQL-Go-ApiServer/graphql/user"
	"mojeico/GraphQL-Go-ApiServer/repository"
	"mojeico/GraphQL-Go-ApiServer/service"
	"net/http"

	"github.com/graphql-go/graphql"
)

var (
	userRepo    = repository.NewUserRepository()
	userService = service.NewUserService(userRepo)
)

func main() {

	documentType := document.NewDocumentTypeGraphql()

	userType := user.NewUserTypeGraphql(documentType)
	userFields := user.NewUserField(userType, userService)
	userMutation := user.NewUserMutation(userType, userService)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: userFields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: userMutation,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

}

// /graphql?query=query+getUser($id:ID){user(id:$id){name}}&variables={"id":"4"}

/*
queryAllUsers := `
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

/*
queryCreate := `
    mutation {
        createUser(UserId: 5, UserName:"test") {
            UserId
			UserName
        }
    }
`
*/
