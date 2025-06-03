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

func (app *App) Run(envs *config.Config) {
	app.DBClient = db.NewMongo(envs.MongoURI, envs.MongoDBName)
	app.RedisClient = db.NewRedis(envs.RedisURI)

	//TODO: КОЛЕКЦИИ ВЫНЕСТИ
	app.Repositories.UserRepository = repository.NewUserRepository(app.DBClient.DB, "users")
	app.Services.UserService = service.NewUserService(app.Repositories.UserRepository)
	app.Handlers.UserHandler = handler.NewUserHandler(app.Services.UserService)

	app.Router = echo.New()
	routes.SetupRoutes(app.Router, *envs, *app.Handlers)
}

func main() {
	var (
		envs = config.LoadConfig()
		app  = &App{}
	)

	app.Run(envs)
	err := app.Router.Start(envs.Port)
	if err != nil {
		log.Printf("failed to start server: %v", err)
	}
}
