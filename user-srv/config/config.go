package config

type ServerConfig struct {
	Name         string             `mapstructure:"server-name"`
	MysqlConfig  MysqlServerConfig  `mapstructure:"mysql"`
	ConsulConfig ConsulServerConfig `mapstructure:"consul"`
}

type MysqlServerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ConsulServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
