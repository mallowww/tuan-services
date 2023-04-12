package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/home", func(c echo.Context) error {
		time.Sleep(time.Millisecond * 10)
		return c.String(http.StatusOK, "ok")
	})
	err := e.Start(":8088")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
