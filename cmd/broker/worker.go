package main

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	apiauth "github.com/amirhosseinf79/renthub_service/internal/services/api_auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/homsa"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jabama"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jajiga"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/mihmansho"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/otaghak"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/logger"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	manager "github.com/amirhosseinf79/renthub_service/internal/services/service_manager"
	"github.com/amirhosseinf79/renthub_service/internal/services/webhook"
)

func main() {
	db := database.NewGormDB(false)
	clientServiceManager := broker.NewClient()

	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)
	requestService := requests.New()

	tokenRepo := persistence.NewTokenRepository(db)
	tokenService := auth.NewTokenService(tokenRepo)
	userRepo := persistence.NewUserRepository(db)
	userService := auth.NewUserService(userRepo, tokenService)
	webhookService := webhook.NewWebhookService(userService, requestService)

	logRepo := persistence.NewLogRepository(db)
	logService := logger.NewLogger(logRepo)

	homsaService := homsa.New(apiAuthService, requestService)
	jabamaService := jabama.New(apiAuthService, requestService)
	jajigaService := jajiga.New(apiAuthService, requestService)
	mihmanshoService := mihmansho.New(apiAuthService, requestService)
	otaghakService := otaghak.New(apiAuthService, requestService)
	shabService := shab.New(apiAuthService, requestService)

	services := map[string]interfaces.ApiService{
		"homsa":     homsaService,
		"jabama":    jabamaService,
		"jajiga":    jajigaService,
		"mihmansho": mihmanshoService,
		"otaghak":   otaghakService,
		"shab":      shabService,
	}

	serviceManager := manager.New(
		services,
		apiAuthService,
		logService,
	)

	fmt.Println("Connecting to worker...")
	broker := broker.NewWorker(
		clientServiceManager,
		serviceManager,
		logService,
		services,
		webhookService,
	)
	broker.StartWorker()
}
