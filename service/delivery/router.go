package delivery

import (
	"github.com/anggardagasta/mit/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"time"
)

type handler struct {
	usersUsecase service.IServiceUsersUsecase
}

func Router(usersUsecase service.IServiceUsersUsecase) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Encoding", "Cookie", "Origin", "X-Api-Key"},
	})
	router.Use(cors.Handler)

	h := handler{usersUsecase: usersUsecase}

	router.Group(func(router chi.Router) {
		router.Route("/v1/users", func(router chi.Router) {
			router.Use(middleware.SetHeader("Content-Type", "application/json"))
			router.Use(middleware.SetHeader("Content-Type", "application/json"))

			router.Post("/register", h.RegisterUser)
			router.Post("/check/phone", h.CheckUserPhone)
			router.Post("/check/email", h.CheckUserEmail)

		})
	})
	return router
}
