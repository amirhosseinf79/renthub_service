package server

func (s server) InitRoutes() {
	api_v1 := s.app.Group("/api/v1")
	api_v2 := s.app.Group("/api/v2")

	s.initAuthRoutes_v1(api_v1)
	s.initManagerRoutes_v1(api_v1)
	s.initApiAuthRoutes_v1(api_v1)
	s.initLoggerRoutes_v1(api_v1)

	s.initManagerRoutes_v2(api_v2)
	s.initApiAuthRoutes_v2(api_v2)
	s.initRecieveRoutes_v2(api_v2)
}
