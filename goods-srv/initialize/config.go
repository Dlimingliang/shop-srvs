package initialize

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	"github.com/Dlimingliang/shop-srvs/goods-srv/config"
	"github.com/Dlimingliang/shop-srvs/goods-srv/global"
)

func InitConfig() {
	configFile := "goods-srv/config-dev.yaml"
	viper.AutomaticEnv()
	isProd := viper.GetBool("IS_PROD")
	if isProd {
		configFile = "goods-srv/config-prod.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Panic(err)
	}

	nacosConfig := config.NacosConfig{}
	err = v.Unmarshal(&nacosConfig)
	if err != nil {
		zap.S().Panic(err)
	}
	zap.S().Infof("nacos配置信息: %v", nacosConfig)

	//加载项目配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   nacosConfig.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		zap.S().Panic(err.Error())
	}

	data, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group})
	if err != nil {
		zap.S().Panic(err)
	}

	err = yaml.Unmarshal([]byte(data), global.ServerConfig)
	if err != nil {
		zap.S().Panic(err)
	}
	zap.S().Infof("goods-srv配置信息: %v", global.ServerConfig)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		zap.S().Error(err)
	}
}
