package api

func (s *Server) routes() {
	s.router.GET("/health", s.handleHealth)
}
