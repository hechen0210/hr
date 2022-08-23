package repository

import "hr/model"

type User struct {
	UserModel model.User
}

func NewUserRepo() *User {
	return &User{}
}
