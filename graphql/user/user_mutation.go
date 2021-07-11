package user

import (
	"github.com/graphql-go/graphql"
	"mojeico/GraphQL-Go-ApiServer/service"
)

func NewUserMutation(userType graphql.Output, userService service.UserService) *graphql.Field {

	createUser := &graphql.Field{
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
		Resolve: userService.CreateUser,
	}

	return createUser
}
