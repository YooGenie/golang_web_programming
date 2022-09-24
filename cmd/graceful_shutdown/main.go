package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println(os.Getpid())
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/short", func(c echo.Context) error { //5초동안 처리하는 API
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	e.GET("/long", func(c echo.Context) error { //30초동안 처리하는 API
		time.Sleep(30 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //시그널이 오면 10초동안 기다려준다.
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}