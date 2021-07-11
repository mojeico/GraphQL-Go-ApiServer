package document

import "github.com/graphql-go/graphql"

func NewDocumentTypeGraphql() *graphql.Object {

	return graphql.NewObject(
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

}
