package main

import (
	"github.com/graphql-go/handler"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {

	fields, mutations := SetUpFieldsAndMutations()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutations,
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
