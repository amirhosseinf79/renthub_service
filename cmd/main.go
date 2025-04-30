package main

import (
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/database"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/persistence"
	"github.com/amirhosseinf79/renthub_service/internal/Infrastructure/server"
	"github.com/amirhosseinf79/renthub_service/internal/application/handler"
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/homsa"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jabama"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/jajiga"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/mihmansho"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/otaghak"
	"github.com/amirhosseinf79/renthub_service/internal/services/api_service/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/auth"
	manager "github.com/amirhosseinf79/renthub_service/internal/services/service_manager"
)

func main() {
	db := database.NewGormDB("host=localhost user= password= dbname= port=5432 sslmode=disable TimeZone=Asia/Tehran", true)

	// User auth system
	authUserService := auth.ImplementAuthUser(db)

	apiRepo := persistence.NewApiAuthRepository(db)
	homsaService := homsa.New(apiRepo)
	jabamaService := jabama.New(apiRepo)
	jajigaService := jajiga.New(apiRepo)
	mihmanshoService := mihmansho.New(apiRepo)
	otaghakService := otaghak.New(apiRepo)
	shabService := shab.New(apiRepo)

	services := map[string]interfaces.ApiService{
		"homsa":     homsaService,
		"jabama":    jabamaService,
		"jajiga":    jajigaService,
		"mihmansho": mihmanshoService,
		"otaghak":   otaghakService,
		"shab":      shabService,
	}

	serviceManager := manager.New(services)

	apiManagerHandler := handler.NewManagerHandler(serviceManager)

	server := server.NewServer(
		authUserService.AuthTokenMiddleware,
		authUserService.UserHandler,
		apiManagerHandler,
	)

	server.InitServer()
	server.InitRoutes()
	server.Start()
}
