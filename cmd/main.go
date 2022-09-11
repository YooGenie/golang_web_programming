package main

import "golang_web_programming/internal"

func main() {
	server := internal.NewDefaultServer()
	server.Run()

	//e:= echo.New()
	//e.GET("/", func(c echo.Context) error { return c.String(c.NoContent(http.StatusOK), "hello world") })
	//e.Logger.Fatal(e.Start(":8080"))
}
