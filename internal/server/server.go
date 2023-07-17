package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ginEngine  *gin.Engine
	httpServer http.Server
}

func NewServer() Server {
	return Server{
		ginEngine: gin.Default(),
	}
}
func (s *Server) ListenAndServe(addr string) error {
	s.httpServer.Addr = addr
	s.httpServer.Handler = s.ginEngine
	return s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
