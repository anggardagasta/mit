package usecase

import (
	"github.com/anggardagasta/mit/models"
	"github.com/anggardagasta/mit/service"
)

type serviceUsersUsecase struct {
	serviceUsersRepo service.IServiceUsersRepository
}

func NewServiceUsersUsecase(serviceUserRepo service.IServiceUsersRepository) service.IServiceUsersUsecase {
	return serviceUsersUsecase{serviceUsersRepo: serviceUserRepo}
}

func (uc serviceUsersUsecase) GetUserByPhoneNumber(input models.FormPhone) (result models.FormUserResponse, err error) {
	result, err = uc.serviceUsersRepo.GetUserByPhoneNumber(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (uc serviceUsersUsecase) GetUserByEmail(input models.FormEmail) (result models.FormUserResponse, err error) {
	result, err = uc.serviceUsersRepo.GetUserByEmail(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (uc serviceUsersUsecase) RegisterUser(input models.FormUserRequest) (err error) {
	err = uc.serviceUsersRepo.InsertUser(input)
	if err != nil {
		return err
	}
	return nil
}
