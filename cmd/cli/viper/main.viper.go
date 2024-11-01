package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port      int `mapstructure:"port"`
		Databases []struct {
			User         string `mapstructure:"user"`
			Password     string `mapstructure:"password"`
			Host         string `mapstructure:"host"`
			DatabaseName string `mapstructure:"database_name"`
		} `mapstructure:"databases"`
	}
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// read config

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// read server config
	fmt.Println("Server Host: ", viper.GetInt("server.port"))
	fmt.Println("Server Host: ", viper.GetString("security.jwt.key"))
}
