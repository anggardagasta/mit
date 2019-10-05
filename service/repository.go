package service

import "github.com/anggardagasta/mit/models"

type IServiceUsersRepository interface {
	GetUserByPhoneNumber(input models.FormPhone) (result models.FormUser, err error)
	GetUserByEmail(input models.FormEmail) (result models.FormUser, err error)
	InsertUser(input models.FormUser) error
}
