package global

import (
	"gorm.io/gorm"

	"github.com/Dlimingliang/shop-srvs/user-srv/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
)
