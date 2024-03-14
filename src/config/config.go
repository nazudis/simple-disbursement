package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read config file: %v", err)
	}
}
