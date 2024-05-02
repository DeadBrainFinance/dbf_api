package utils

import (
	"log"

	"github.com/spf13/viper"
)

type EnvConfigs struct {
    POSTGRES_DB string
    POSTGRES_USER string
    POSTGRES_PASSWORD string
    PGDATA string
    POSTGRES_PORT string
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

