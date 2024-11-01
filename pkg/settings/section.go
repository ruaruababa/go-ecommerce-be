package settings

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Mysql  MysqlSettings  `mapstructure:"mysql"`
	Logger LoggerSettings `mapstructure:"logger"`
}

type MysqlSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
	TimeOut  int    `mapstructure:"connection_timeout"`
}

type LoggerSettings struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	File       string `mapstructure:"file"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
