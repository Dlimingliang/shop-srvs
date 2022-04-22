package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID         int32     `gorm:"primarykey"`
	CreateUser int32     `gorm:"column:create_user"`
	UpdateUser int32     `gorm:"column:update_user"`
	CreatedAt  time.Time `gorm:"column:create_time;type:datetime"`
	UpdatedAt  int64     `gorm:"column:update_time;autoUpdateTime"`
	IsDelete   bool      `gorm:"not null;default:0;comment:0:enabled 1:disabled"`
}
