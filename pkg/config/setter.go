package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the config")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		os.Exit(1)
	}

}
