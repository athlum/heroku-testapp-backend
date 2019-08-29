package server

import (
	"github.com/labstack/echo"
)

type Server struct{}

func (s *Server) Register(g *echo.Group) {
	sg := g.Group("api")

	sg.GET("/query", s.Query)
	sg.GET("/assets", s.QueryAssets)
}

func (s *Server) Query(c echo.Context) error {
	return c.String(200, "query")
}

func (s *Server) QueryAssets(c echo.Context) error {
	return c.String(200, "queryAssets")
}
