package server

func (s server) InitRoutes() {
	api := s.app.Group("/api/v1")

	// s.initAuthRoutes(api)
	api.Use(s.tokenMiddleware.CheckTokenAuth)
}

// func (s server) initAuthRoutes(api fiber.Router) {
// 	api.Post("/auth/register", s.userHandler.RegisterUser)
// 	api.Post("/auth/login", s.userHandler.LoginUser)
// }
