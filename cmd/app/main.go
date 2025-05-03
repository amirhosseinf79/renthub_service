package main

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	"github.com/amirhosseinf79/renthub_service/internal/application/handler"
	"github.com/amirhosseinf79/renthub_service/internal/application/middleware"
	apiauth "github.com/amirhosseinf79/renthub_service/internal/services/api_auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
)

func main() {
	db := database.NewGormDB(false)
	clientServiceManager := broker.NewClient()

	// User auth system
	authUserService := auth.ImplementAuthUser(db)

	// api auth model
	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)

	apiManagerValidator := middleware.NewValidator()
	apiTokenMiddleware := middleware.NewApiTokenMiddleware(clientServiceManager, apiAuthService)
	apiManagerHandler := handler.NewManagerHandler(clientServiceManager, apiAuthService)

	server := server.NewServer(
		authUserService.AuthTokenMiddleware,
		authUserService.UserHandler,
		apiTokenMiddleware,
		apiManagerValidator,
		apiManagerHandler,
	)

	server.InitServer()
	server.InitRoutes()
	server.Start()
}
