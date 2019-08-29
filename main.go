package main

import (
	"github.com/labstack/echo"
	_ "net/http/pprof"

	"github.com/athlum/heroku-testapp-backend/server"
)

func main() {
	e := echo.New()

	root := e.Group("/")
	new(server.Server).Register(root)

	e.Start("0.0.0.0:80")
}
