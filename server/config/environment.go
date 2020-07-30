package config

import (
	"github.com/spf13/viper"
	"log"
)

func ConfigEnvVariables() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}


}