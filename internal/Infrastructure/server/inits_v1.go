package server

import "github.com/gofiber/fiber/v3"

func (s server) initAuthRoutes_v1(api fiber.Router) {
	api.Post("/auth/register", s.userHandler.RegisterUser)
	api.Post("/auth/login", s.userHandler.LoginUser)
	api.Post("/auth/refresh", s.userHandler.RefreshToken)
	api.Post("/auth/update", s.userHandler.UpdateUser, s.tokenMiddleware.CheckTokenAuth)
}

func (s server) initManagerRoutes_v1(api fiber.Router) {
	update := api.Group(
		"/service/update",
		s.tokenMiddleware.CheckTokenAuth,
		s.apiTokenMiddleware.ApiAuthValidator,
		s.apiManagerValidator.DateCheck,
	)
	update.Post("/calendar", s.apiManagerHandler.UpdateCalendar, s.apiManagerValidator.CalendarUpdate)
	update.Post("/discount", s.apiManagerHandler.UpdateDiscount, s.apiManagerValidator.DiscountUpdate)
	update.Post("/reservation", s.apiManagerHandler.UpdateMinNight, s.apiManagerValidator.MinNightUpdate)
	update.Post("/price", s.apiManagerHandler.UpdatePrice, s.apiManagerValidator.PriceUpdate)
}

func (s server) initApiAuthRoutes_v1(api fiber.Router) {
	auth := api.Group("/service/auth", s.tokenMiddleware.CheckTokenAuth)
	auth.Post("/check", s.apiManagerHandler.CheckAuth, s.apiTokenMiddleware.ApiAuthValidator)
	auth.Post("/refresh", s.apiManagerHandler.RefreshToken, s.apiManagerValidator.RefReshTokenCheck)
	auth.Post("/send-otp", s.apiManagerHandler.SendServiceOTP, s.apiManagerValidator.SendOTPCheck)
	auth.Post("/verify-otp", s.apiManagerHandler.VerifyServiceOTP, s.apiManagerValidator.VerifyOTPCheck)
	auth.Post("/sign-out", s.apiManagerHandler.SignOutClient)
	// auth.Post("/", s.apiManagerHandler.TokenLogin)
}
