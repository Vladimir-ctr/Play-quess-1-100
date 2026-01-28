package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1

	const maxAttempts = 10

	var inputNumber int

	var quess bool = false

	fmt.Printf("Я загадал число от 1 до 100. У тебя %d попыток чтобы отгадать! \n", maxAttempts)
	fmt.Println("Введите число")

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Попытка %d: введите число: ", attempt)
		_, err := fmt.Scan(&inputNumber)
		if err != nil {
			fmt.Println("Пожалуйста введите корректное число")
			attempt--
			continue
		}

		if inputNumber < randomNumber {
			fmt.Println("Неверно, секретное число больше!")
		} else if inputNumber > randomNumber {
			fmt.Println("Неверно, секретное число меньше!")
		} else if inputNumber == randomNumber {
			fmt.Println("Ура!!! Вы угадали число!")
			quess = true
			break
		}
	}

	if !quess {
		fmt.Println("Попытки кончились, вы проиграли")
		fmt.Printf("Загаданное число было: %d\n", randomNumber)
	}

}
