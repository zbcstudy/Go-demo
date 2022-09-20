package main

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	// 添加中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (b bool, err error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 {
			b = true
			err = nil
			return
		}
		if username == "joe" || username == "zbc" {
			return true, nil
		}
		return false, nil
	}))

	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			println("request to users")
			return next(context)
		}
	}

	e.GET("/users", func(context echo.Context) error {
		return context.String(http.StatusOK, "/users")
	}, track)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	// body
	e.PUT("/users", func(context echo.Context) error {
		user := new(User)
		if err := context.Bind(user); err != nil {
			return err
		}
		return context.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":8090"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+",member:"+member)

}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+",email:"+email)
}
