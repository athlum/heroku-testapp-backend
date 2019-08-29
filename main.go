package main

import (
	"github.com/labstack/echo"
	_ "net/http/pprof"
	"os"

	"github.com/athlum/heroku-testapp-backend/server"
)

var port string

func init() {
	port = os.Getenv("PORT")
}

func main() {
	e := echo.New()

	root := e.Group("/")
	new(server.Server).Register(root)

	if err := e.Start("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
