package delivery

import (
	"github.com/anggardagasta/mit/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handler struct {
	usersUsecase service.IServiceUsersUsecase
}

func Router(usersUsecase service.IServiceUsersUsecase) *chi.Mux {
	router := chi.NewRouter()

	h := handler{usersUsecase: usersUsecase}

	router.Group(func(router chi.Router) {
		router.Route("/v1/users", func(router chi.Router) {
			router.Use(middleware.SetHeader("Content-Type", "application/json"))

			router.Post("/register", h.RegisterUser)
			router.Post("/check/phone", h.CheckUserPhone)
			router.Post("/check/email", h.CheckUserEmail)

		})
	})
	return router
}
