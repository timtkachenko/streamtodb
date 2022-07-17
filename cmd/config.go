package main

import (
	"fmt"
	"github.com/google/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file. Using default config...")
	}
	_ = viper.BindEnv("DEBUG", "DEBUG")
	// db
	_ = viper.BindEnv("DB_HOSTNAME", "DB_HOSTNAME")
	_ = viper.BindEnv("DB_USERNAME", "DB_USERNAME")
	_ = viper.BindEnv("POSTGRES_PASSWORD", "POSTGRES_PASSWORD")
	_ = viper.BindEnv("DB_DATABASE", "DB_DATABASE")

	viper.Set("POSTGRES_DSN", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", viper.Get("DB_HOSTNAME"), viper.Get("DB_USERNAME"),
		viper.Get("DB_DATABASE"), viper.Get("POSTGRES_PASSWORD")))
	viper.AutomaticEnv()
}
