package main

import (
	"basicproject/initializers"
	transactiontransport "basicproject/internal/module/transaction/transport"
	"basicproject/internal/module/user/transport"
	tokenfactory "basicproject/internal/token_factory"
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
	tokenFactory tokenfactory.TokenFactory
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

	service := &Service{
		tokenFactory: tokenfactory.CreateTokenFactory(
			[]byte("secret-key"),
			jwt.SigningMethodHS256,
		),
	}

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/sign-up", usertransport.HandleCreateUser(initializers.DB))
		user.POST("/sign-in", usertransport.HandleSignIn(initializers.DB, service.tokenFactory))
	}

	transaction := v1.Group("/transaction")
	{
		transaction.POST("/create", transactiontransport.HandleCreateTransaction(initializers.DB))
	}

	e.Logger.Fatal(e.Start(":1323"))
}
