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
	"github.com/amirhosseinf79/renthub_service/internal/services/logger"
	manager "github.com/amirhosseinf79/renthub_service/internal/services/service_manager"
)

func main() {
	db := database.NewGormDB(false)
	clientServiceManager := broker.NewClient()

	apiRepo := persistence.NewApiAuthRepository(db)
	apiAuthService := apiauth.NewApiAuthService(apiRepo)

	logRepo := persistence.NewLogRepository(db)
	logService := logger.NewLogger(logRepo)

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
	)

	fmt.Println("Connecting to worker...")
	broker := broker.NewWorker(
		clientServiceManager,
		serviceManager,
		logService,
	)
	broker.StartWorker()
}
