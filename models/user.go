package models

type User struct {
	UserId       int      `json:"user_id"`
	UserName     string   `json:"user_name"`
	UserPassword string   `json:"user_password"`
	UserDocument Document `json:"user_document"`
}
