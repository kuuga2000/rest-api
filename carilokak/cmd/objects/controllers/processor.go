package controllers

import (
	"carilokak/cmd/objects/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Objects struct {
	ObjectId     int    `json:"id"`
	ObjectName   string `json:"name"`
	ObjectGender int    `json:"gender"`
}

// e.POST("/object/create", objects.CreateObject)
func CreateObject(c echo.Context) error {
	newobject := models.Objects{}
	c.Bind(&newobject)
	newObject, err := models.CreateObject(newobject)
	//newObject := new(Objects)
	//newObject, err := models.CreateObject(newobject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newObject)
}
