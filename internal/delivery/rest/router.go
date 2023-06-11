package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"logity/internal/delivery/rest/handlers/auth"
	"logity/internal/delivery/rest/handlers/log"
	"logity/internal/delivery/rest/handlers/operating"
	"logity/internal/delivery/rest/handlers/room"
	"logity/internal/domain/usecase"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

func RegisterRouting(r chi.Router, env *usecase.Env) {

	authHandler := auth.NewHandler(env.AuthUsecase, env.RoomUsecase)
	roomHandler := room.NewHandler(env.RoomUsecase)
	operatingHandler := operating.NewHandler(env.OperatingUsecase)
	logHandler := log.NewHandler(env.LogUsecase)

	//not secure routes
	r.Group(func(r chi.Router) {
		authHandler.Register(r)
	})

	//secure routes
	r.Group(func(r chi.Router) {
		r.Use(authHandler.AuthMiddleware)
		roomHandler.Register(r)
		operatingHandler.Register(r)
		logHandler.Register(r)
	})
}
