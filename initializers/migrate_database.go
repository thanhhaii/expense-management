package initializers

import (
	transactionmodel "basicproject/internal/module/transaction/model"
	"basicproject/internal/module/user/model"
)

func AutoSyncDatabase() {
	err := DB.AutoMigrate(&usermodel.User{}, &transactionmodel.Transaction{})
	if err != nil {
		panic("Failed when migrate table")
	}
}
