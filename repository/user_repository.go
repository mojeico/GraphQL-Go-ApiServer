package repository

import "mojeico/GraphQL-Go-ApiServer/models"

var user1 models.User = models.User{
	UserId:       1,
	UserName:     "testUserName1",
	UserPassword: "testPass1",
	UserDocument: models.Document{
		DocumentId:      1,
		DocumentName:    "PassportName1",
		DocumentNumber:  23521,
		DocumentExpDate: "12.01.2024",
		UserId:          1,
	},
}

var user2 models.User = models.User{
	UserId:       2,
	UserName:     "testUserName2",
	UserPassword: "testPass2",
	UserDocument: models.Document{
		DocumentId:      5,
		DocumentName:    "PassportName3",
		DocumentNumber:  836502,
		DocumentExpDate: "13.11.2094",
		UserId:          2,
	},
}

var user3 models.User = models.User{
	UserId:       3,
	UserName:     "testUserName3",
	UserPassword: "testPass3",
	UserDocument: models.Document{
		DocumentId:      1,
		DocumentName:    "PassportName3",
		DocumentNumber:  04731514,
		DocumentExpDate: "27.02.2004",
		UserId:          3,
	},
}

var users = []models.User{user1, user2, user3}

type UserRepository interface {
	CreateUser(user models.User) models.User
	UpdateUser(userId int, user models.User)
	DeleteUser(userId int)

	GetAllUser() []models.User
	GetUserById(userId int) models.User
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u userRepository) CreateUser(user models.User) models.User {
	users = append(users, user)

	return user
}

func (u userRepository) UpdateUser(userId int, user models.User) {

	newUsers := make([]models.User, 0)

	for _, v := range users {
		if v.UserId == userId {
			newUsers = append(newUsers, user)
		}
	}

	users = newUsers
}

func (u userRepository) DeleteUser(userId int) {

	newUsers := make([]models.User, 0)

	for _, v := range users {
		if v.UserId != userId {
			newUsers = append(newUsers, v)
		}
	}

	users = newUsers
}

func (u userRepository) GetAllUser() []models.User {
	return users
}

func (u userRepository) GetUserById(userId int) models.User {

	for _, v := range users {
		if v.UserId == userId {
			return v
		}
	}

	return models.User{}

}
