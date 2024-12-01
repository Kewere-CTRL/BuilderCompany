package main

import (
	"fmt"
	"github.com/Kewere-CTRL/BuilderCompany/internal/db"
	"github.com/Kewere-CTRL/BuilderCompany/internal/router"
	"log"
	"net"

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

func getLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Ошибка получения интерфейсов: ", err)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal("Ошибка получения адресов интерфейса: ", err)
		}

		for _, addr := range addrs {
			// Ищем IPv4 адрес
			if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "IP не найден"
}

func main() {
	initConfig()
	db.InitDB()

	serverPort := viper.GetInt("server.port")

	// Получаем локальный IP
	localIP := getLocalIP()
	fmt.Printf("Сервер запущен на адресе: http://%s:%d\n", localIP, serverPort)

	server := router.SetupRouter()
	server.Run(fmt.Sprintf("%s:%d", localIP, serverPort))
}
