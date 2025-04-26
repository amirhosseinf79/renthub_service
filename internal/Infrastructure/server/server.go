package server

import (
	"log"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type server struct {
	app             *fiber.App
	tokenMiddleware interfaces.TokenMiddleware
	userHandler     interfaces.UserHandler
}

func NewServer(
	tokenMiddleware interfaces.TokenMiddleware,
	userHandler interfaces.UserHandler,
) *server {
	return &server{
		tokenMiddleware: tokenMiddleware,
		userHandler:     userHandler,
	}
}

func (s *server) InitServer() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))
	s.app = app
}

func (s *server) Start() {
	err := s.app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
