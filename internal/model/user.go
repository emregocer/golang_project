package model

import "time"

type User struct {
	Base
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func NewUserModel(id int, username string, password string, email string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		Base:     Base{Id: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		Username: username,
		Password: password,
		Email:    email,
	}
}
