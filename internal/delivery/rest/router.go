package rest

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"logity/config"
	"logity/docs"
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

func ConfigureSwagger(cfg *config.App) {
	docs.SwaggerInfo.Title = "Logity"
	docs.SwaggerInfo.Description = "rest api logity"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", cfg.Host, cfg.ApiPort)
}

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RegisterRouting(r chi.Router, env *usecase.Env, cfg *config.App) {
	ConfigureSwagger(cfg)
	r.Mount("/swagger", httpSwagger.WrapHandler)

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
