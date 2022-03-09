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
	UserName string     `gorm:"column:user_name;type:varchar(20);not null;unique;index:index_user_name; comment '用户名'"`
	Mobile   string     `gorm:"column:mobile;type:varchar(11);not null;unique;index:index_mobile; comment '电话'"`
	Password string     `gorm:"column:password;type:varchar(100);not null; comment '密码'"`
	Gender   bool       `gorm:"column:gender;type:tinyint(2); comment '性别 1:男 2:女'"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime; comment '生日'"`
	Role     int        `gorm:"column:role;type:tinyint(2);not null;default:2; comment '角色 1:管理员 2:普通用户'"`
}
