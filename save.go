package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func saveResults(win bool, attempts int) {
	// Создаем обьект результата
	object := GameResults{
		Date:     time.Now().Format("2000-01-01 00:00:00"), // прлучаем текущее время и дату и форматируем его в удобный нам формат
		Result:   "win",
		Attempts: attempts,
	}

	if !win {
		object.Result = "lose"
	}

	var results []GameResults
	data, err := os.ReadFile("results.json") // встроенная функция читает результаты если они были, возвращает ошибку например если файл не найден
	if err == nil {
		json.Unmarshal(data, &results) // Если ошибок нет, тогда распаковываем JSON (встроенная функция превращает JSON текст в переменные ГО с которыми можно раьотать) в переменную results
	}

	results = append(results, object)

	jsonData, err := json.MarshalIndent(results, "", "   ") // Формируем JSON
	if err != nil {
		fmt.Println("Ошибка подготовки данных", err)
		return
	}

	err = os.WriteFile("results.json", jsonData, 0644) // Записываем рещультаты в файл
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
	} else {
		fmt.Println("Результат игры сохранён в results.json")
	}
}
