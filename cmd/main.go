package main

import (
	"os"

	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	"github.com/amirhosseinf79/renthub_service/internal/application/handler"
	"github.com/amirhosseinf79/renthub_service/internal/application/middleware"
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	apiauth "github.com/amirhosseinf79/renthub_service/internal/services/api_auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/homsa"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jabama"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jajiga"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/mihmansho"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/otaghak"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
	"github.com/amirhosseinf79/renthub_service/internal/services/logger"
	manager "github.com/amirhosseinf79/renthub_service/internal/services/service_manager"
)

func main() {
	dbConfig := os.Getenv("DB")
	db := database.NewGormDB(dbConfig, false)

	// User auth system
	authUserService := auth.ImplementAuthUser(db)

	apiRepo := persistence.NewApiAuthRepository(db)
	logRepo := persistence.NewLogRepository(db)

	logService := logger.NewLogger(logRepo)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)

	homsaService := homsa.New(apiAuthService)
	jabamaService := jabama.New(apiAuthService)
	jajigaService := jajiga.New(apiAuthService)
	mihmanshoService := mihmansho.New(apiAuthService)
	otaghakService := otaghak.New(apiAuthService)
	shabService := shab.New(apiAuthService)

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
		dto.ManagerConfig{
			SendWebHookSeperately: false,
		})

	apiManagerValidator := middleware.NewValidator()
	apiTokenMiddleware := middleware.NewApiTokenMiddleware(apiAuthService)
	apiManagerHandler := handler.NewManagerHandler(serviceManager, apiAuthService)

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
