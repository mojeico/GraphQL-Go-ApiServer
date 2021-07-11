package service

import (
	"github.com/graphql-go/graphql"
	"mojeico/GraphQL-Go-ApiServer/models"
	"mojeico/GraphQL-Go-ApiServer/repository"
)

type UserService interface {
	CreateUser() func(p graphql.ResolveParams) (interface{}, error)
	UpdateUser() func(p graphql.ResolveParams) (interface{}, error)
	DeleteUser() func(p graphql.ResolveParams) (interface{}, error)

	GetAllUser() func(p graphql.ResolveParams) (interface{}, error)
	GetUserById() func(p graphql.ResolveParams) (interface{}, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u userService) CreateUser() func(p graphql.ResolveParams) (interface{}, error) {

	return func(params graphql.ResolveParams) (interface{}, error) {
		user := models.User{
			UserId:   params.Args["UserId"].(int),
			UserName: params.Args["UserName"].(string),
			//UserPassword: params.Args["UserPassword"].(string),
			//UserDocument: params.Args["UserDocument"].(models.Document),
		}

		userCreated := u.repo.CreateUser(user)
		return userCreated, nil

	}
}

func (u userService) UpdateUser() func(p graphql.ResolveParams) (interface{}, error) {

	return func(p graphql.ResolveParams) (interface{}, error) {
		userId := p.Args["id"].(int)
		//u.repo.UpdateUser(userId, user)
		return userId, nil
	}
}

func (u userService) DeleteUser() func(p graphql.ResolveParams) (interface{}, error) {

	return func(p graphql.ResolveParams) (interface{}, error) {
		userId := p.Args["id"].(int)
		u.repo.DeleteUser(userId)

		return userId, nil
	}

}

func (u userService) GetAllUser() func(p graphql.ResolveParams) (interface{}, error) {

	return func(p graphql.ResolveParams) (interface{}, error) {
		users := u.repo.GetAllUser()
		return users, nil
	}

}

func (u userService) GetUserById() func(p graphql.ResolveParams) (interface{}, error) {

	return func(p graphql.ResolveParams) (interface{}, error) {

		userId := p.Args["id"].(int)
		return u.repo.GetUserById(userId), nil

	}

}
