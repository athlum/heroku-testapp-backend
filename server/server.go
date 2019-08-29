package server

import (
	"fmt"
	"github.com/labstack/echo"
	"os"
	"time"

	"github.com/athlum/heroku-testapp-backend/model"
)

type Server struct{}

func (s *Server) Register(g *echo.Group) {
	sg := g.Group("api")

	sg.GET("/query", s.Query)
	sg.GET("/insert", s.Insert)
	sg.GET("/assets", s.QueryAssets)
	sg.GET("/config", s.Config)
}

func (s *Server) Insert(c echo.Context) error {
	u := &model.User{Name: fmt.Sprintf("%v", time.Now())}
	if err := model.DB().Insert(u); err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, fmt.Sprintf("insert id: %v", u.ID))
}

func (s *Server) Query(c echo.Context) error {
	v, err := model.DB().SelectInt("select count(id) from user")
	if err != nil {
		return c.String(500, err.Error())
	}
	r := fmt.Sprintf("%v rows\n", v)
	us := []*model.User{}
	_, err = model.DB().Select(&us, "select * from user")
	if err != nil {
		return c.String(500, err.Error())
	}
	for _, u := range us {
		r += fmt.Sprintf("Id: %v - Name: %v\n", u.ID, u.Name)
	}
	return c.String(200, r)
}

func (s *Server) QueryAssets(c echo.Context) error {
	return c.String(200, "queryAssets")
}

func (s *Server) Config(c echo.Context) error {
	return c.JSONBlob(200, []byte(os.Getenv("CONFIG")))
}
