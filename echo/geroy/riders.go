package geroy

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Rider struct {
    Name  string `json:"name" form:"name" query:"name"`
    Skill string `json:"skill" form:"skill" query:"skill"`
}

func createRider(c echo.Context) error {
	rider := new(Rider)
    if err := c.Bind(rider); err != nil {
        return err
    }

    return c.String(http.StatusOK, "Asd")
}