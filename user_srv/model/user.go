package model

import (
	"time"
)

type BaseModel struct {
	ID         int32     `gorm:"primarykey"`
	CreateUser int32     `gorm:"column: create_user"`
	UpdateUser int32     `gorm:"column: update_user"`
	CreateTime time.Time `gorm:"column: create_time"`
	UpdateTime time.Time `gorm:"column: update_time"`
	IsDelete   bool
}

type User struct {
	BaseModel
	UserName string `gorm:"column: user_name;type:varchar(20);unique;index:index_user_name"`
}
