package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/home", func(c echo.Context) error {
		time.Sleep(time.Millisecond * 10)
		return c.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":1677"))
}
