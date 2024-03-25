package main

import (
	"WebServerUbuntu/internal/ConfigServer"
	"WebServerUbuntu/internal/HandlerServer"
	dirServe "WebServerUbuntu/pkg/FunctionsString"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/braintree/manners"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func main() {

	log.Println("Trying to read a config")

	dir, _ := os.Getwd() // Получение директории файла

	dirFile := dirServe.GetDir(dir, 2) // Получаем директорию конфига

	dirFile += "internal/ConfigServer/configServer.toml"

	config := ConfigServer.ConfigServer{}

	_, err := toml.DecodeFile(dirFile, &config)
	if err != nil {
		log.Fatal("Fail to open file!")
	}

	// Подготвока к запуску сервера
	ch := make(chan os.Signal, 1) // Создаем сигнал

	signal.Notify(ch, os.Interrupt, os.Kill)

	go ListenForShutdown(ch)

	router := HandlerServer.SetHandlerHTTPRouter()

	HandlerServer.HandlerFunc(router)

	// Запуск сервера
	log.Println("Trying start server")
	errServer := manners.ListenAndServe(":"+strconv.Itoa(int(config.Port)), router) // Создаем сервер
	if errors.Is(errServer, http.ErrServerClosed) {
		log.Fatal("Failed to start server")
	}
}

// ListenForShutdown Функция для безопасного закрытия подключения к серверу
func ListenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
	log.Println("Closing the server")
}
