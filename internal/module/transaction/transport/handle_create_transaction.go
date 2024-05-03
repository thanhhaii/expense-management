package transactiontransport

import (
	"basicproject/internal/common"
	transactionbiz "basicproject/internal/module/transaction/business"
	transactionmodel "basicproject/internal/module/transaction/model"
	transactionstorage "basicproject/internal/module/transaction/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func HandleCreateTransaction(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			transactionItem transactionmodel.TransactionCreate
			err             error
		)

		if err = c.Bind(&transactionItem); err != nil {
			return c.JSON(http.StatusBadRequest, common.Response[bool]{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
				Success: false,
			})
		}

		storage := transactionstorage.NewPostgreStorage(db)
		createTransactionBiz := transactionbiz.NewCreateTransactionBiz(storage)
		transactionCode, err := createTransactionBiz.CreateNewTransaction(c.Request().Context(), &transactionItem)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.Response[bool]{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
				Success: false,
			})
		}

		return c.JSON(http.StatusOK, common.Response[string]{
			Data:    transactionCode,
			Status:  http.StatusOK,
			Success: false,
		})
	}
}
