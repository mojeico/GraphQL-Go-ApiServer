package main

import (
	"github.com/graphql-go/graphql"
	"mojeico/GraphQL-Go-ApiServer/graphql/document"
	"mojeico/GraphQL-Go-ApiServer/graphql/user"
	"mojeico/GraphQL-Go-ApiServer/repository"
	"mojeico/GraphQL-Go-ApiServer/service"
)

var (
	documentType = document.NewDocumentTypeGraphql()

	userRepo                          = repository.NewUserRepository()
	userService                       = service.NewUserService(userRepo)
	userType                          = user.NewUserTypeGraphql(documentType)
	getUserById, userList, deleteUser = user.NewUserField(userType, userService)
	createUser                        = user.NewUserMutation(userType, userService)
)

func SetUpFieldsAndMutations() (graphql.Fields, *graphql.Object) {

	queryFields := graphql.Fields{
		"getUserById": getUserById,
		"userList":    userList,
		"deleteUser":  deleteUser,
	}

	mutations := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createUser": createUser,
			},
		},
	)

	return queryFields, mutations
}
