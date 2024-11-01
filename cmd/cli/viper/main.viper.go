package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Databases []struct {
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Host         string `mapstructure:"host"`
		DatabaseName string `mapstructure:"database_name"`
	} `mapstructure:"databases"`
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

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	for _, db := range config.Databases {
		fmt.Println("Database User: ", db.User)
		fmt.Println("Database Password: ", db.Password)
		fmt.Println("Database Host: ", db.Host)
		fmt.Println("Database Name: ", db.DatabaseName)
	}
}
