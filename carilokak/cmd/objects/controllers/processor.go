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

	message := struct {
		Success bool           `json:"success"`
		Data    models.Objects `json:"data"`
	}{
		Success: true,
		Data:    newObject,
	}

	return c.JSON(http.StatusCreated, message)
}

func DeleteObject(c echo.Context) error {
	objectId := c.Param("id")
	id := models.DeleteObject(objectId)
	message := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: id + " has been deleted successfully",
	}
	return c.JSON(http.StatusOK, message)
}

func GetData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
