package model

import (
	"time"
)

type BaseModel struct {
	ID         int32     `gorm:"primarykey"`
	CreateUser string    `gorm:"column:create_user"`
	UpdateUser string    `gorm:"column:update_user"`
	CreatedAt  time.Time `gorm:"column:create_time"`
	UpdatedAt  time.Time `gorm:"column:update_time"`
	IsDelete   bool      `gorm:"not null;default:0;comment:0:enabled 1:disabled"`
}

type User struct {
	BaseModel
	UserName string     `gorm:"column:user_name;type:varchar(20);not null;comment:用户名"`
	Mobile   string     `gorm:"column:mobile;type:varchar(11);not null;unique;index:index_mobile;comment:电话"`
	Password string     `gorm:"column:password;type:varchar(100);not null; comment:密码"`
	Gender   string     `gorm:"column:gender;type:varchar(10); comment:性别 男 女"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime; comment:生日"`
	Role     int8       `gorm:"column:role;;not null;default:2; comment:角色 1:管理员 2:普通用户"`
}
