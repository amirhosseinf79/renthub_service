package main

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
)

func main() {
	db := database.NewGormDB("host=localhost user= password= dbname= port=5432 sslmode=disable TimeZone=Asia/Tehran", true)

	// User auth system
	authUserService := auth.ImplementAuthUser(db)

	server := server.NewServer(
		authUserService.AuthTokenMiddleware,
		authUserService.UserHandler,
	)

	server.InitServer()
	server.InitRoutes()
	server.Start()
}
