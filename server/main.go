package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()

	e.GET("/healthcheck", healthCheck)

	e.Logger.Fatal(e.Start(":1323"))
}