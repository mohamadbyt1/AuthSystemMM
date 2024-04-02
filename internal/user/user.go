package user

import (
_"github.com/go-playground/validator/v10"
)

type User struct {
	Id       string 
	Username string 
	Password string 
	Role string 
}
type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}