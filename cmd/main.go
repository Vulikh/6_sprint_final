package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)

	httpServer := server.New(logger)

	logger.Printf("Сервер запущен")
	if err := httpServer.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
