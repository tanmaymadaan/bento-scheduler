package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tanmaymadaan/bento-scheduler/pkg/database"
)

type Server struct {
	router *gin.Engine
	db     *database.CassandraSession
}

func NewServer(db *database.CassandraSession) *Server {
	s := &Server{
		router: gin.Default(),
		db:     db,
	}
	s.routes()
	return s
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
