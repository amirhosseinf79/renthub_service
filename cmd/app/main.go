package main

import (
	broker_v1 "github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker/v1"
	broker_v2 "github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker/v2"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	handler_v1 "github.com/amirhosseinf79/renthub_service/internal/application/handler/v1"
	handler_v2 "github.com/amirhosseinf79/renthub_service/internal/application/handler/v2"
	middleware_v1 "github.com/amirhosseinf79/renthub_service/internal/application/middleware/v1"
	middleware_v2 "github.com/amirhosseinf79/renthub_service/internal/application/middleware/v2"
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
)

func main() {
	db := database.NewGormDB(false)
	broker_v1 := broker_v1.NewClient()
	broker_v2 := broker_v2.NewClient()

	// User auth system
	authUserService_v1 := auth_v1.ImplementAuthUser(db)

	loggerRepo := persistence.NewLogRepository(db)
	loggerService := logger.NewLogger(loggerRepo)

	// api auth model
	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)
	requestService := requests.New()

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

	// manager middlewares & handler
	apiManagerValidator_v1 := middleware_v1.NewValidator()
	apiTokenMiddleware_v1 := middleware_v1.NewApiTokenMiddleware(broker_v1, apiAuthService)
	apiManagerHandler_v1 := handler_v1.NewManagerHandler(services, broker_v1, apiAuthService)
	loggerHandler := handler_v1.NewLogHandler(loggerService)

	// manager middlewares & handler
	apiManagerValidator_v2 := middleware_v2.NewValidator()
	apiTokenMiddleware_v2 := middleware_v2.NewApiTokenMiddleware(broker_v2, apiAuthService)
	apiManagerHandler_v2 := handler_v2.NewManagerHandler(services, broker_v2, apiAuthService)

	server := server.NewServer(
		authUserService_v1.AuthTokenMiddleware,
		authUserService_v1.UserHandler,
		loggerHandler,
		apiManagerValidator_v1,
		apiTokenMiddleware_v1,
		apiManagerHandler_v1,
		apiManagerValidator_v2,
		apiTokenMiddleware_v2,
		apiManagerHandler_v2,
	)

	server.InitServer()
	server.InitRoutes()
	server.Start()
}
