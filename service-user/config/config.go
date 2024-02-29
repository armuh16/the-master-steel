package config

import "github.com/spf13/viper"

var c *Config

const (
	Development = "Development"
	Production  = "Production"
)

type Config struct {
	Env      string `yaml:"env"`
	Postgres struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslMode"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbName"`
	} `yaml:"postgres"`
}

func Get() *Config {
	return c
}

func SetConfig() {

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err.Error())
	}

	viper.WatchConfig()
}
