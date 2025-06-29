package auth

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/application/handler"
	"github.com/amirhosseinf79/renthub_service/internal/application/middleware"
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"gorm.io/gorm"
)

type authUserService struct {
	UserHandler         interfaces.UserHandler
	AuthTokenMiddleware interfaces.TokenMiddleware
}

func ImplementAuthUser(db *gorm.DB) *authUserService {
	userRepo := persistence.NewUserRepository(db)
	tokenRepo := persistence.NewTokenRepository(db)
	tokenService := NewTokenService(tokenRepo)
	userservice := NewUserService(userRepo, tokenService)
	authTokenMiddleware := middleware.NewAuthTokenMiddleware(tokenService)
	userHandler := handler.NewUserHandler(userservice, tokenService)

	return &authUserService{
		UserHandler:         userHandler,
		AuthTokenMiddleware: authTokenMiddleware,
	}
}
