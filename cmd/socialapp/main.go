package main

import (
	"basicproject/initializers"
	usertransport "basicproject/internal/user/transport"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.AutoSyncDatabase()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/create-user", usertransport.HandleCreateUser(initializers.DB))
	}

	e.Logger.Fatal(e.Start(":1323"))
}
