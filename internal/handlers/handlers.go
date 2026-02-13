package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	// Поддерживаем только GET
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, req, "../index.html")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	// Поддерживаем только POST
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusInternalServerError)
		return
	}
	// Парсим форму
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Ошибка парсинга формы %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Выбираем файл из формы
	file, _, err := req.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка получения файла %v", err)
		http.Error(w, "Ошибка получения файла ", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Читаем данные из файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения", http.StatusInternalServerError)
		return
	}
	// Конвертируем данные файла
	result, err := service.Convert(string(content))
	if err != nil {
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}
	// Создаем локальный файл и записываем в него
	safeTime := time.Now().UTC().String()
	ext := filepath.Ext(".txt")
	fileName := fmt.Sprintf("%s%s", safeTime, ext)
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, result)
}
