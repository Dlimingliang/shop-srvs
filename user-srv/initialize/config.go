package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/Dlimingliang/shop_srvs/user-srv/global"
)

func InitConfig() {
	configFile := "user-srv/config-dev.yaml"
	viper.AutomaticEnv()
	isProd := viper.GetBool("IS_PROD")
	if isProd {
		configFile = "user-web/config-prod.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Panic(err)
	}

	err = v.Unmarshal(global.ServerConfig)
	if err != nil {
		zap.S().Panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		err = v.ReadInConfig()
		if err != nil {
			zap.S().Errorw("文件改变,重新读取Config失败", "msg", err.Error())
		}
		err = v.Unmarshal(global.ServerConfig)
		if err != nil {
			zap.S().Errorw("文件改变,重新转换Config失败", "msg", err.Error())
		}
	})
}
