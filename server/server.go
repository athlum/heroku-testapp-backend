package server

import (
	"fmt"
	"github.com/labstack/echo"
	"os"

	"github.com/athlum/heroku-testapp-backend/model"
)

type Server struct{}

func (s *Server) Register(g *echo.Group) {
	sg := g.Group("api")

	sg.GET("/query", s.Query)
	sg.GET("/assets", s.QueryAssets)
	sg.GET("/config", s.Config)
}

func (s *Server) Query(c echo.Context) error {
	v, err := model.DB().SelectInt("select count(id) from user")
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, fmt.Sprintf("%v rows", v))
}

func (s *Server) QueryAssets(c echo.Context) error {
	return c.String(200, "queryAssets")
}

func (s *Server) Config(c echo.Context) error {
	return c.JSONBlob(200, []byte(os.Getenv("CONFIG")))
}
