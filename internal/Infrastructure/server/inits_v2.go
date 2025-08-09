package server

import "github.com/gofiber/fiber/v3"

func (s server) initManagerRoutes_v2(api fiber.Router) {
	update := api.Group(
		"/service/update",
		s.tokenMiddleware.CheckTokenAuth,
		s.apiManagerValidator_v2.DateCheck,
	)
	update.Post("/calendar", s.apiManagerHandler_v2.UpdateCalendar, s.apiManagerValidator_v2.CalendarUpdate)
	update.Post("/discount", s.apiManagerHandler_v2.UpdateDiscount, s.apiManagerValidator_v2.DiscountUpdate)
	update.Post("/reservation", s.apiManagerHandler_v2.UpdateMinNight, s.apiManagerValidator_v2.MinNightUpdate)
	update.Post("/price", s.apiManagerHandler_v2.UpdatePrice, s.apiManagerValidator_v2.PriceUpdate)
}

func (s server) initApiAuthRoutes_v2(api fiber.Router) {
	auth := api.Group("/service/auth", s.tokenMiddleware.CheckTokenAuth)
	auth.Post("/check", s.apiManagerHandler_v2.CheckAuth, s.apiTokenMiddleware_v2.ApiAuthValidator)
	auth.Post("/refresh", s.apiManagerHandler_v2.RefreshToken, s.apiManagerValidator_v2.RefReshTokenCheck)
	auth.Post("/send-otp", s.apiManagerHandler_v2.SendServiceOTP, s.apiManagerValidator_v2.SendOTPCheck)
	auth.Post("/verify-otp", s.apiManagerHandler_v2.VerifyServiceOTP, s.apiManagerValidator_v2.VerifyOTPCheck)
	auth.Post("/sign-out", s.apiManagerHandler_v2.SignOutClient, s.apiManagerValidator_v2.SignOutValidator)
	// auth.Post("/", s.apiManagerHandler.TokenLogin)
}
