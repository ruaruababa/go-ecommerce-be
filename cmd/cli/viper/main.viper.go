package main

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
}
