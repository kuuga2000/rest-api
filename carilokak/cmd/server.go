package main

import (
	objectController "carilokak/cmd/objects/controllers"
	"carilokak/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/object/create", objectController.CreateObject)
	e.DELETE("/object/delete/:id", objectController.DeleteObject)
	e.GET("/object/getdata", objectController.GetData)

	configs.InitDB()
	e.Logger.Fatal(e.Start(":1323"))
}
