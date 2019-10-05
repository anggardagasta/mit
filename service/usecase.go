package service

import "github.com/anggardagasta/mit/models"

type IServiceUsersUsecase interface {
	GetUserByPhoneNumber(input models.FormPhone) (result models.FormUser, err error)
	GetUserByEmail(input models.FormEmail) (result models.FormUser, err error)
	RegisterUser(input models.FormUser) error
}
