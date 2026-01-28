package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxAttempts = 10

var quess bool = false

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	var inputNumber int

	fmt.Printf("Я загадал число от 1 до 100. У тебя %d попыток чтобы отгадать! \n", maxAttempts)
	fmt.Println("Введите число")

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Попытка %d: введите число: ", attempt)
		// Проверка ввода
		_, err := fmt.Scan(&inputNumber)
		if err != nil {
			fmt.Println("Пожалуйста введите корректное число")
			fmt.Scanf("%s") // Очистка буфера
			attempt--
			continue
		}

		if hints(inputNumber, randomNumber) {
			return
		}
	}

	fmt.Println("Попытки кончились, вы проиграли")
	fmt.Printf("Загаданное число было: %d\n", randomNumber)

}

// пишем дополнительную функцию abs для расчета разницы по модулю (т.е функция возвращает число всегда положительное)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hints(inputNumber int, randomNumber int) bool {
	if inputNumber == randomNumber {
		fmt.Println("Ура!!! Вы угадали секретное число")
		return true
	}

	if inputNumber <= randomNumber {
		fmt.Println("Секретное число больше")
	} else {
		fmt.Println("Секретное число меньше")
	}

	// создаем переменную в которой будем вычислять разницу между секретным числом и входным и передавать это число в функцию abs
	different := abs(randomNumber - inputNumber)
	// прописываем условия для подсказок
	if different <= 5 {
		fmt.Println("Горячо")
	} else if different <= 15 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Холодно")
	}
	return false

}
