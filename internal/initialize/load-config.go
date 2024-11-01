package initialize

import (
	"fmt"

	"go-ecommerce-be/globals"

	"github.com/spf13/viper"
)

func InitLoadConfig() {
	// load config
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// read config

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&globals.Config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// for _, db := range config.Databases {
	// 	fmt.Println("Database User: ", db.User)
	// 	fmt.Println("Database Password: ", db.Password)
	// 	fmt.Println("Database Host: ", db.Host)
	// 	fmt.Println("Database Name: ", db.DatabaseName)
	// }
}
