package auth_v1

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	handler_v1 "github.com/amirhosseinf79/renthub_service/internal/application/handler/v1"
	middleware_v1 "github.com/amirhosseinf79/renthub_service/internal/application/middleware/v1"
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
	authTokenMiddleware := middleware_v1.NewAuthTokenMiddleware(tokenService)
	userHandler := handler_v1.NewUserHandler(userservice, tokenService)

	return &authUserService{
		UserHandler:         userHandler,
		AuthTokenMiddleware: authTokenMiddleware,
	}
}
