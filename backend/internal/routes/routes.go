package routes

import (
	"github.com/JuDyas/buy-sell-platform/backend/config"
	"github.com/JuDyas/buy-sell-platform/backend/internal/app"
	"github.com/JuDyas/buy-sell-platform/backend/internal/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRoutes(e *echo.Echo, envs config.Config, handlers app.Handlers) {
	var (
		v1    = e.Group("/api/v1")
		users = v1.Group("/users")
	)

	e.Use(middleware.AuthMiddleware(envs.JWTSecret))
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/api/v1/docs")
	})
	users.POST("/register", handlers.UserHandler.Register(envs.JWTSecret))
	users.POST("/login", handlers.UserHandler.Login(envs.JWTSecret))
	users.PUT("", handlers.UserHandler.Update())
	users.GET("/:id", handlers.UserHandler.GetByID())
	users.POST("/upload-avatar", handlers.UserHandler.UploadAvatar())
}
