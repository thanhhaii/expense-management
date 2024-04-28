package userstorage

import "gorm.io/gorm"

type postgreStorage struct {
	db *gorm.DB
}

func NewPostgreStorage(db *gorm.DB) *postgreStorage {
	return &postgreStorage{db: db}
}
