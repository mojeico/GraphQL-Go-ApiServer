package service

import (
	"github.com/graphql-go/graphql"
	"mojeico/GraphQL-Go-ApiServer/models"
	"mojeico/GraphQL-Go-ApiServer/repository"
)

type UserService interface {
	CreateUser(p graphql.ResolveParams) (interface{}, error)
	UpdateUser(p graphql.ResolveParams) (interface{}, error)
	DeleteUser(p graphql.ResolveParams) (interface{}, error)

	GetAllUser(p graphql.ResolveParams) (interface{}, error)
	GetUserById(p graphql.ResolveParams) (interface{}, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u userService) CreateUser(p graphql.ResolveParams) (interface{}, error) {

	user := models.User{
		UserId:   p.Args["UserId"].(int),
		UserName: p.Args["UserName"].(string),
		//UserPassword: params.Args["UserPassword"].(string),
		//UserDocument: params.Args["UserDocument"].(models.Document),
	}

	userCreated := u.repo.CreateUser(user)
	return userCreated, nil
}

func (u userService) UpdateUser(p graphql.ResolveParams) (interface{}, error) {

	userId := p.Args["id"].(int)
	//u.repo.UpdateUser(userId, user)
	return userId, nil
}

func (u userService) DeleteUser(p graphql.ResolveParams) (interface{}, error) {

	userId := p.Args["id"].(int)
	u.repo.DeleteUser(userId)
	return userId, nil
}

func (u userService) GetAllUser(p graphql.ResolveParams) (interface{}, error) {

	users := u.repo.GetAllUser()
	return users, nil
}

func (u userService) GetUserById(p graphql.ResolveParams) (interface{}, error) {

	userId := p.Args["id"].(int)
	return u.repo.GetUserById(userId), nil
}
