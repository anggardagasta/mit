package service

import "github.com/anggardagasta/mit/models"

type IServiceUsersUsecase interface {
	GetUserByPhoneNumber(input models.FormPhone) (result models.FormUserResponse, err error)
	GetUserByEmail(input models.FormEmail) (result models.FormUserResponse, err error)
	RegisterUser(input models.FormUserRequest) error
}
