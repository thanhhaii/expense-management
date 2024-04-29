package usertransport

import (
	"basicproject/internal/common"
	userbiz "basicproject/internal/module/user/business"
	usermodel "basicproject/internal/module/user/model"
	userstorage "basicproject/internal/module/user/storage"
	tokenfactory "basicproject/internal/token_factory"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func HandleSignIn(db *gorm.DB, tokenService tokenfactory.TokenFactory) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payloadSignUp usermodel.SignInModel
			err           error
			token         string
		)

		if err = c.Bind(&payloadSignUp); err != nil {
			return c.JSON(http.StatusBadRequest, common.Response[bool]{
				Success: false,
				Status:  400,
				Message: err.Error(),
			})
		}

		storage := userstorage.NewPostgreStorage(db)
		business := userbiz.CreateSignInBiz(storage, tokenService)

		token, err = business.SignIn(c.Request().Context(), &payloadSignUp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.Response[bool]{
				Success: false,
				Status:  500,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, common.Response[string]{
			Success: true,
			Status:  200,
			Data:    token,
		})
	}
}
