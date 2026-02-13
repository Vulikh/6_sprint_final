package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {

	// Удаляем пробелы и переносы строк
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("файл пуст или содержит только пробелы")
	}

	// Проверяем наличие символов азбуки Морзе
	isMorse := true
	for _, r := range data {
		if r != '.' && r != '-' && r != ' ' && r != '\n' {
			isMorse = false
			break
		}
	}

	result := ""

	if isMorse {
		// Конвертируем из Морзе в текст
		result = morse.ToText(data)
		if result == "" {
			return "", errors.New("не удалось распознать азбуку Морзе")
		}
	} else {
		// Иначе конвертируем из текста в Морзе
		result = morse.ToMorse(data)
		if result == "" {
			return "", errors.New("Ошибка конвертации текста в азбуку Морзе")
		}
	}

	return result, nil

}
