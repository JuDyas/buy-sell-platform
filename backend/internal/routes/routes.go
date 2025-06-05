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
		v1         = e.Group("/api/v1")
		admin      = v1.Group("/admin", middleware.AdminMiddleware)
		users      = v1.Group("/users")
		adds       = v1.Group("/adds")
		categories = v1.Group("/categories")
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
	users.GET("/:id/adds", handlers.AdvertHandler.GetByAuthor())

	adds.POST("", handlers.AdvertHandler.Create())
	adds.PUT("/:id", handlers.AdvertHandler.Update())
	adds.GET("/:id", handlers.AdvertHandler.GetByID())
	adds.DELETE("/:id", handlers.AdvertHandler.SoftDelete())
	adds.POST("/upload-images", handlers.AdvertHandler.UploadImages())
	adds.GET("", handlers.AdvertHandler.GetAll())

	categories.GET("", handlers.CategoryHandler.GetAll())
	categories.GET("/:id", handlers.CategoryHandler.GetByID())
	categories.PUT("/:id", handlers.CategoryHandler.Update())
	categories.GET("/:id/adds", handlers.AdvertHandler.GetByCategory())

	admin.POST("/categories", handlers.CategoryHandler.Create())
	admin.PUT("/categories/:id", handlers.CategoryHandler.Update())
	admin.DELETE("/categories/:id", handlers.CategoryHandler.Delete())
}
