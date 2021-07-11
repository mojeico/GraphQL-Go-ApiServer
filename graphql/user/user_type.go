package user

import "github.com/graphql-go/graphql"

func NewUserTypeGraphql(documentField graphql.Output) *graphql.Object {

	return graphql.NewObject(
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
					Type: documentField,
				},
			},
			IsTypeOf:    nil,
			Description: "My User",
		},
	)

}
