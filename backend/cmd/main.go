package main

import (
	"github.com/JuDyas/buy-sell-platform/backend/config"
	"github.com/JuDyas/buy-sell-platform/backend/internal/app"
	"github.com/JuDyas/buy-sell-platform/backend/internal/db"
	"github.com/JuDyas/buy-sell-platform/backend/internal/handler"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/JuDyas/buy-sell-platform/backend/internal/routes"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"log"
)

type App struct {
	Router       *echo.Echo
	DBClient     *db.Mongo
	RedisClient  *redis.Client
	Handlers     *app.Handlers
	Services     *app.Service
	Repositories *app.Repository
}

func (a *App) Run(envs *config.Config) {
	a.DBClient = db.NewMongo(envs.MongoURI, envs.MongoDBName)
	a.RedisClient = db.NewRedis(envs.RedisURI)

	a.Repositories = &app.Repository{}
	a.Services = &app.Service{}
	a.Handlers = &app.Handlers{}

	//TODO: КОЛЕКЦИИ ВЫНЕСТИ
	a.Repositories.UserRepository = repository.NewUserRepository(a.DBClient.DB, "users")
	a.Services.UserService = service.NewUserService(a.Repositories.UserRepository)
	a.Handlers.UserHandler = handler.NewUserHandler(a.Services.UserService)

	a.Repositories.AdvertRepository = repository.NewAdvertRepository(a.DBClient.DB, "adds")
	a.Services.AdvertService = service.NewAdvertService(a.Repositories.AdvertRepository)
	a.Handlers.AdvertHandler = handler.NewAdvertHandler(a.Services.AdvertService)

	a.Router = echo.New()
	a.Router.Static("/static", "static")
	routes.SetupRoutes(a.Router, *envs, *a.Handlers)
}

func main() {
	var (
		envs = config.LoadConfig()
		a    = &App{}
	)

	a.Run(envs)
	err := a.Router.Start(":" + envs.Port)
	if err != nil {
		log.Printf("failed to start server: %v", err)
	}
}
