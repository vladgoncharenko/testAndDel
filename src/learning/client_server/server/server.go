package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"syscall"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		fmt.Println(syscall.EPIPE)
		time.Sleep(30 * time.Second)
		return context.JSON(http.StatusOK, struct{ Status string }{Status: "ok"})
	})

	e.Logger.Fatal(e.Start(":9897"))
}
