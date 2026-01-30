// функция вывод подсказок
package main

import (
	"fmt"
)

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
