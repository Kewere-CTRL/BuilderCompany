package main

import (
	"fmt"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/router"
	"log"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфига: %s", err)
	}
	viper.AutomaticEnv()
}

func main() {

	initConfig()
	db.InitDB()

	serverPort := viper.GetInt("server.port")

	server := router.SetupRouter()
	server.Run(fmt.Sprintf("0.0.0.0:%d", serverPort))
}
