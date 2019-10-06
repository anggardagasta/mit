package service

import "github.com/anggardagasta/mit/models"

type IServiceUsersRepository interface {
	GetUserByPhoneNumber(input models.FormPhone) (result models.FormUserResponse, err error)
	GetUserByEmail(input models.FormEmail) (result models.FormUserResponse, err error)
	InsertUser(input models.FormUserRequest) error
}
