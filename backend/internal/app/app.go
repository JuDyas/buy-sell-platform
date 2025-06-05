package app

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/handler"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
)

type Handlers struct {
	UserHandler     *handler.UserHandler
	AdvertHandler   *handler.AdvertHandler
	CategoryHandler *handler.CategoryHandler
}

type Service struct {
	UserService     service.UserService
	AdvertService   service.AdvertService
	CategoryService service.CategoriesService
}

type Repository struct {
	UserRepository     repository.UserRepository
	AdvertRepository   repository.AdvertRepository
	CategoryRepository repository.CategoryRepository
}
