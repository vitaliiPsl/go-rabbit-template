package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port                      int    `mapstructure:"PORT"`
	RABBIT_URL                string `mapstructure:"RABBIT_URL"`
	RABBIT_MAIN_QUEUE_NAME    string `mapstructure:"RABBIT_MAIN_QUEUE_NAME"`
	RABBIT_MAIN_EXCHANGE_NAME string `mapstructure:"RABBIT_MAIN_EXCHANGE_NAME"`
	RABBIT_MAIN_EXCHANGE_TYPE string `mapstructure:"RABBIT_MAIN_EXCHANGE_TYPE"`
	RABBIT_MAIN_ROUTING_KEY   string `mapstructure:"RABBIT_MAIN_ROUTING_KEY"`
}

func ReadConfig() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Can't find config file: %v\n", err.Error())
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v", err)
		panic(err)
	}

	return &config
}
