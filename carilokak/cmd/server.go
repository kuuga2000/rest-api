package main

import (
	"carilokak/cmd/objects"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/object/create", objects.CreateObject)
	e.POST("/object/getall", objects.GetAllObjects)
	e.POST("/object/detail/:id", objects.CreateObject)
	e.POST("/object/edit/:id", objects.EditObject)
	e.POST("/object/delete/:id", objects.CreateObject)
	e.POST("/object/message", objects.Message)
	e.Logger.Fatal(e.Start(":1323"))
}
