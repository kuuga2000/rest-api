package main

import (
	"net/http"
	"sync"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Job string `json:"job"`
	}
)

type News struct {
    Title  string `json:"title" form:"title" query:"title"`
    Author string `json:"author" form:"author" query:"author"`
}

var (
	users = map[int]*user{}
	seq   = 1
	lock  = sync.Mutex{}
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/show", showUser)
	e.GET("/name/:id", getUser)
	e.POST("/create", createUser)
	e.POST("/createPost", createPost)
	e.GET("/testx", createRider)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.JSON(http.StatusOK, "name:" + id)
}

func createRider(c echo.Context) error {
	return c.JSON(http.StatusOK, "asd:asd")
}

func showUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &user {
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusOK, u)
}

func createPost(c echo.Context) error {
	news := new(News)
    if err := c.Bind(news); err != nil {
        return err
    }

    return c.JSON(http.StatusOK, news)
}