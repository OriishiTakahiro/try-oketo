package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	podIP = ""
)

func main() {
	podIP = os.Getenv("POD_IP")

	// echo instance
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/ip", func(c echo.Context) error {
		return c.String(http.StatusOK, podIP)
	})

	// start Server
	e.Logger.Fatal(e.Start(":80"))
}
