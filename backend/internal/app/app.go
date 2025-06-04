package app

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/handler"
	"github.com/JuDyas/buy-sell-platform/backend/internal/repository"
	"github.com/JuDyas/buy-sell-platform/backend/internal/service"
)

type Handlers struct {
	UserHandler *handler.UserHandler
}

type Service struct {
	UserService service.UserService
}

type Repository struct {
	UserRepository repository.UserRepository
}
