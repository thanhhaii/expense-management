package main

import (
	"basicproject/initializers"
	tokenfactory "basicproject/internal/token_factory"
	usertransport "basicproject/internal/user/transport"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.AutoSyncDatabase()
}

type Service struct {
	TokenFactory tokenfactory.TokenFactory
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte("secret-key"),
		SigningMethod: echojwt.AlgorithmHS256,
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return strings.Contains(path, "/sign-in") || strings.Contains(path, "/sign-up")
		},
	}))

	_ = &Service{
		TokenFactory: tokenfactory.CreateTokenFactory(
			[]byte("secret-key"),
			jwt.SigningMethodES256,
		),
	}

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/create-user", usertransport.HandleCreateUser(initializers.DB))
	}

	e.Logger.Fatal(e.Start(":1323"))
}
