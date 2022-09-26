package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// "net/http"
)

func main(){
	//echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	//start server
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error{
	return c.String(http.StatusOK, "Hello, World!")
}