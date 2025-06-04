package app

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/handler"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
)

type Handlers struct {
	UserHandler   *handler.UserHandler
	AdvertHandler *handler.AdvertHandler
}

type Service struct {
	UserService   service.UserService
	AdvertService service.AdvertService
}

type Repository struct {
	UserRepository   repository.UserRepository
	AdvertRepository repository.AdvertRepository
}
