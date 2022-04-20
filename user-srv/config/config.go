package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}

type ServerConfig struct {
	Name         string             `yaml:"server-name"`
	MysqlConfig  MysqlServerConfig  `yaml:"mysql"`
	ConsulConfig ConsulServerConfig `yaml:"consul"`
}

type MysqlServerConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ConsulServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
