package main

import (
	"fmt"
)

// Ведение файла results.json, где сохраняется дата, исход (победа/проигрыш), количество попыток.
// После типа поля указывается поле в json
type GameResults struct {
	Date     string `json:"date"`     // Дата и время в формате "ГГГГ-ММ-ДД ЧЧ:ММ:СС"
	Result   string `json:"outcome"`  // "win" — победа, "lose" — проигрыш
	Attempts int    `json:"attempts"` // Сколько попыток использовано
}

var maxAttempts int
var maxNumber int

func main() {
	fmt.Println("Добро пожаловать в игру Угадай число!")
	fmt.Println("Выберите уровень: Easy, Medium, Hard")
	levelChoice()

	// Запускаем игру — получаем результат
	win, attempts := game(maxAttempts, maxNumber)

	// Сохраняем: дата, исход, попытки
	saveResults(win, attempts)

	// Спрашиваем, хочет ли игрок сыграть ещё
	retry()
}

// пишем дополнительную функцию abs для расчета разницы по модулю (т.е функция возвращает число всегда положительное)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
