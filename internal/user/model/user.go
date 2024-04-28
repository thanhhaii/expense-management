package usermodel

import "time"

type User struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Email     string     `json:"email" gorm:"column:email;index:idx_email,unique"`
	Password  string     `json:"password" gorm:"column:password;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (User) TableName() string {
	return "user"
}
