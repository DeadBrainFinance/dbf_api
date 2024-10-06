package utils

import (
	"log"

	"github.com/spf13/viper"
)

type EnvConfigs struct {
    DB string
    DB_USER string
    DB_PASSWORD string
    DB_PORT string
    HOST string
}

func(config *EnvConfigs) LoadEnvVariables() {
    viper.AddConfigPath(".")
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatal("Error reading env file", err)
    }

    if err := viper.Unmarshal(&config); err != nil {
        log.Fatal(err)
    }
    return
}

