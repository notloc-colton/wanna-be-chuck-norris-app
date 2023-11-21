// TODO: Stubbed server package that will need to be hardened and unit tests added
package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttachRoute interface {
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}
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
func (s *Server) AttachRoute() AttachRoute {
	return s.ginEngine
}
