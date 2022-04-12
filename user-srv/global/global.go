package global

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/shop_srvs/user-srv/config"
)

const DbIp = "127.0.0.1"

var (
	ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
)
