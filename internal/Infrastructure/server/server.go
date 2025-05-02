package server

import (
	"log"
	"os"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type server struct {
	app                 *fiber.App
	tokenMiddleware     interfaces.TokenMiddleware
	userHandler         interfaces.UserHandler
	apiManagerValidator interfaces.ValidatorInterface
	apiTokenMiddleware  interfaces.ApiAuthMiddleware
	apiManagerHandler   interfaces.ManagerHandlerInterface
}

func NewServer(
	tokenMiddleware interfaces.TokenMiddleware,
	userHandler interfaces.UserHandler,
	apiTokenMiddleware interfaces.ApiAuthMiddleware,
	apiManagerValidator interfaces.ValidatorInterface,
	apiManagerHandler interfaces.ManagerHandlerInterface,
) *server {
	return &server{
		tokenMiddleware:     tokenMiddleware,
		userHandler:         userHandler,
		apiManagerValidator: apiManagerValidator,
		apiTokenMiddleware:  apiTokenMiddleware,
		apiManagerHandler:   apiManagerHandler,
	}
}

func (s *server) InitServer() {
	app := fiber.New(fiber.Config{
		TrustProxy: true,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies: []string{
				"127.0.0.1",
			},
		},
	})
	app.Use(logger.New(logger.Config{}))
	s.app = app
}

func (s *server) Start() {
	envPort := os.Getenv("PORT")
	err := s.app.Listen("127.0.0.1:"+envPort, fiber.ListenConfig{EnablePrefork: false})
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
