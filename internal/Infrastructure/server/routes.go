package server

import "github.com/gofiber/fiber/v3"

func (s server) InitRoutes() {
	api := s.app.Group("/api/v1")

	// s.initAuthRoutes(api)
	s.initManagerRoutes(api)
	s.initApiAuthRouts(api)
}

// func (s server) initAuthRoutes(api fiber.Router) {
// 	api.Post("/auth/register", s.userHandler.RegisterUser)
// 	api.Post("/auth/login", s.userHandler.LoginUser)
// }

func (s server) initManagerRoutes(api fiber.Router) {
	update := api.Group("/service/update", s.tokenMiddleware.CheckTokenAuth, s.apiManagerValidator.DateCheck)
	update.Post("/calendar", s.apiManagerHandler.UpdateCalendar, s.apiManagerValidator.CalendarUpdate)
	update.Post("/discount", s.apiManagerHandler.UpdateDiscount, s.apiManagerValidator.DiscountUpdate)
	update.Post("/reservation", s.apiManagerHandler.UpdateMinNight, s.apiManagerValidator.MinNightUpdate)
	update.Post("/price", s.apiManagerHandler.UpdatePrice, s.apiManagerValidator.PriceUpdate)
}

func (s server) initApiAuthRouts(api fiber.Router) {
	auth := api.Group("/service/auth", s.tokenMiddleware.CheckTokenAuth)
	auth.Post("/", s.apiManagerHandler.TokenLogin, s.tokenMiddleware.CheckTokenAuth)
	auth.Post("/check", s.apiManagerHandler.CheckAuth)
}
