package main

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/broker"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	"github.com/amirhosseinf79/renthub_service/internal/application/handler"
	"github.com/amirhosseinf79/renthub_service/internal/application/middleware"
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	apiauth "github.com/amirhosseinf79/renthub_service/internal/services/api_auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/homsa"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jabama"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jajiga"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/mihmansho"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/otaghak"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
)

func main() {
	db := database.NewGormDB(false)
	clientServiceManager := broker.NewClient()

	// User auth system
	authUserService := auth.ImplementAuthUser(db)

	// api auth model
	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)
	requestService := requests.New()

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

	// manager middlewares & handler
	apiManagerValidator := middleware.NewValidator()
	apiTokenMiddleware := middleware.NewApiTokenMiddleware(clientServiceManager, apiAuthService)
	apiManagerHandler := handler.NewManagerHandler(services, clientServiceManager, apiAuthService)

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
