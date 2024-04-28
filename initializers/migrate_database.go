package initializers

import usermodel "basicproject/internal/user/model"

func AutoSyncDatabase() {
	err := DB.AutoMigrate(&usermodel.User{})
	if err != nil {
		panic("Failed when migrate table")
	}
}
