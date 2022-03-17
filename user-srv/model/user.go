package model

import (
	"time"
)

type BaseModel struct {
	ID         uint64    `gorm:"primarykey"`
	CreateUser uint64    `gorm:"column:create_user"`
	UpdateUser uint64    `gorm:"column:update_user"`
	CreatedAt  time.Time `gorm:"column:create_time;type:datetime"`
	UpdatedAt  int       `gorm:"column:update_time;autoUpdateTime"`
	IsDelete   bool      `gorm:"not null;default:0;comment:0:enabled 1:disabled"`
}

const (
	GenderUnKnow = 0
	GenderMale   = 1
	GenderFemale = 2
)

type User struct {
	BaseModel
	UserName string     `gorm:"column:user_name;type:varchar(20);not null;comment:用户名"`
	Mobile   string     `gorm:"column:mobile;type:varchar(11);not null;unique;index:index_mobile;comment:电话"`
	Password string     `gorm:"column:password;type:varchar(100);not null; comment:密码"`
	Gender   uint8      `gorm:"column:gender; comment:性别 0:未知 1:男 2:女"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime; comment:生日"`
	Role     uint8      `gorm:"column:role;;not null;default:2; comment:角色 1:管理员 2:普通用户"`
}
