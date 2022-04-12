package global

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/shop_srvs/user-srv/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
)
