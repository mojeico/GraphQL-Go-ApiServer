package user

import (
	"github.com/graphql-go/graphql"
	"mojeico/GraphQL-Go-ApiServer/service"
)

func NewUserField(userType graphql.Output, userService service.UserService) (*graphql.Field, *graphql.Field, *graphql.Field) {

	getUserById := &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: userService.GetUserById(),
	}

	userList := &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "Get User List",
		Resolve:     userService.GetAllUser(),
	}

	deleteUser := &graphql.Field{
		Type: graphql.Int,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve:     userService.DeleteUser(),
		Description: "Delete User by id",
	}

	return getUserById, userList, deleteUser

}
