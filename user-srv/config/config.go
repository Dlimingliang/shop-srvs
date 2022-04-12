package config

type ServerConfig struct {
	MysqlConfig MysqlServerConfig `mapstructure:"mysql"`
}

type MysqlServerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
