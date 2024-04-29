package initializers

import (
	"basicproject/internal/module/user/model"
)

func AutoSyncDatabase() {
	err := DB.AutoMigrate(&usermodel.User{})
	if err != nil {
		panic("Failed when migrate table")
	}
}
