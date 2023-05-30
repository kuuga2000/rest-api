package main

import (
	objectController "carilokak/cmd/objects/controllers"
	"carilokak/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST(configs.Version()+"/object/create", objectController.CreateObject)
	e.GET(configs.Version()+"/object/list", objectController.GetObjects)
	e.PUT(configs.Version()+"/object/edit", objectController.UpdateObject)
	e.DELETE(configs.Version()+"/object/delete/:id", objectController.DeleteObject)
	e.GET(configs.Version()+"/object/getdata", objectController.GetData)

	configs.InitDB()
	e.Logger.Fatal(e.Start(":1323"))
}
