package config

import (
	"github.com/anggardagasta/mit/service/delivery"
	"github.com/anggardagasta/mit/service/repository/mysql"
	"github.com/anggardagasta/mit/service/usecase"
)

func (cfg *Config) InitService() error {
	serviceUserRepo := mysql.NewServiceUsersRepository(cfg.DB)
	serviceUserUsecase := usecase.NewServiceUsersUsecase(serviceUserRepo)

	cfg.Route = delivery.Router(serviceUserUsecase)

	return nil
}
