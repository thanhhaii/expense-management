package usertransport

import (
	userbiz "basicproject/internal/user/business"
	usermodel "basicproject/internal/user/model"
	userstorage "basicproject/internal/user/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func HandleCreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userItem usermodel.User

		if err := c.Bind(&userItem); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   err.Error(),
				"success": false,
			})
		}

		storage := userstorage.NewPostgreStorage(db)
		createUserBiz := userbiz.NewCreateUserBiz(storage)

		if err := createUserBiz.CreateUser(c.Request().Context(), &userItem); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   err.Error(),
				"success": false,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    userItem.Id,
			"success": true,
		})
	}
}
