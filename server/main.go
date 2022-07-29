package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lil-shimon/lil-pubsub.git/server/websocket"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()

	e.GET("/healthcheck", healthCheck)
	e.GET("/ws/:topic", websocket.ServeWs)
	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":1323"))
}