package config

import "github.com/spf13/viper"

type AppConfig struct {
	Database databaseConfig `yaml:"database"`
}

type databaseConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

var Config AppConfig

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
