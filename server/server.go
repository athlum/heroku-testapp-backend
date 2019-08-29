package server

import (
	"github.com/labstack/echo"
	"os"
)

type Server struct{}

func (s *Server) Register(g *echo.Group) {
	sg := g.Group("api")

	sg.GET("/query", s.Query)
	sg.GET("/assets", s.QueryAssets)
	sg.GET("/config", s.Config)
}

func (s *Server) Query(c echo.Context) error {
	return c.String(200, "query")
}

func (s *Server) QueryAssets(c echo.Context) error {
	return c.String(200, "queryAssets")
}

func (s *Server) Config(c echo.Context) error {
	return c.JSONBlob(200, []byte(os.Getenv("CONFIG")))
}
