package server

import (
	"log"
	"os"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type server struct {
	app                    *fiber.App
	tokenMiddleware        interfaces.TokenMiddleware
	userHandler            interfaces.UserHandler
	loggerHnadler          interfaces.LoggerHandler
	apiManagerValidator    interfaces.ValidatorInterface
	apiTokenMiddleware     interfaces.ApiAuthMiddleware
	apiManagerHandler      interfaces.ManagerHandlerInterface
	apiManagerValidator_v2 interfaces.ValidatorInterface_v2
	apiTokenMiddleware_v2  interfaces.ApiAuthMiddleware
	apiManagerHandler_v2   interfaces.ManagerHandlerInterface
	apiRecieveManager_v2   interfaces.RecieveHandler
}

func NewServer(
	tokenMiddleware interfaces.TokenMiddleware,
	userHandler interfaces.UserHandler,
	loggerHnadler interfaces.LoggerHandler,
	apiManagerValidator interfaces.ValidatorInterface,
	apiTokenMiddleware interfaces.ApiAuthMiddleware,
	apiManagerHandler interfaces.ManagerHandlerInterface,
	apiManagerValidator_v2 interfaces.ValidatorInterface_v2,
	apiTokenMiddleware_v2 interfaces.ApiAuthMiddleware,
	apiManagerHandler_v2 interfaces.ManagerHandlerInterface,
	apiRecieveManager_v2 interfaces.RecieveHandler,
) *server {
	return &server{
		tokenMiddleware:        tokenMiddleware,
		userHandler:            userHandler,
		loggerHnadler:          loggerHnadler,
		apiManagerValidator:    apiManagerValidator,
		apiTokenMiddleware:     apiTokenMiddleware,
		apiManagerHandler:      apiManagerHandler,
		apiManagerValidator_v2: apiManagerValidator_v2,
		apiTokenMiddleware_v2:  apiTokenMiddleware_v2,
		apiManagerHandler_v2:   apiManagerHandler_v2,
		apiRecieveManager_v2:   apiRecieveManager_v2,
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
	err := s.app.Listen(":"+envPort, fiber.ListenConfig{EnablePrefork: false})
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
