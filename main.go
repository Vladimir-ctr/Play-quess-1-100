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

func hints(inputNumber int, randomNumber int) bool {
	if inputNumber < randomNumber {
		fmt.Println("Неверно, секретное число больше!")
	} else if inputNumber > randomNumber {
		fmt.Println("Неверно, секретное число меньше!")
	} else if inputNumber == randomNumber {
		fmt.Println("Ура!!! Вы угадали число!")
		return true
	}
	return false
}
