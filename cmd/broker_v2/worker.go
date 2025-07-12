package main

import (
	"fmt"

	broker_v2 "github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker/v2"
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
	auth_v1 "github.com/amirhosseinf79/renthub_service/internal/services/auth/v1"
	"github.com/amirhosseinf79/renthub_service/internal/services/chromium"
	"github.com/amirhosseinf79/renthub_service/internal/services/logger"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	manager_v2 "github.com/amirhosseinf79/renthub_service/internal/services/service_manager/v2"
	webhook_v2 "github.com/amirhosseinf79/renthub_service/internal/services/webhook/v2"
)

func main() {
	db := database.NewGormDB(false)
	clientServiceManager := broker_v2.NewClient()

	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)
	requestService := requests.New()

	tokenRepo := persistence.NewTokenRepository(db)
	userRepo := persistence.NewUserRepository(db)

	tokenService := auth_v1.NewTokenService(tokenRepo)
	userService := auth_v1.NewUserService(userRepo, tokenService)

	webhookService := webhook_v2.NewWebhookService(userService, requestService)

	logRepo := persistence.NewLogRepository(db)
	logService := logger.NewLogger(logRepo)

	chromiumService := chromium.NewChromiumService()
	defer chromiumService.Close()

	homsaService := homsa.New(apiAuthService, requestService)
	jabamaService := jabama.New(apiAuthService, requestService)
	jajigaService := jajiga.New(apiAuthService, requestService)
	mihmanshoService := mihmansho.New(apiAuthService, requestService, chromiumService)
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

	serviceManager := manager_v2.New(
		services,
		apiAuthService,
		logService,
	)

	fmt.Println("Connecting to worker...")
	broker := broker_v2.NewWorker(
		clientServiceManager,
		serviceManager,
		logService,
		services,
		webhookService,
	)
	broker.StartWorker()
}
