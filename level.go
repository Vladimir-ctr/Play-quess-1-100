// функция выбора уровня сложности
package main

import (
	"fmt"
	"strings"
)

func levelChoice() {
	fmt.Println("Выберите уровень: Easy, Medium, Hard")
	var inputLevel string

	_, err := fmt.Scan(&inputLevel) // так работает fmt.Scan - сканирует поле ввода если ожидаемый тип неверный выдает ошибку
	if err == nil {
		inputLevel = strings.ToLower(inputLevel) //Привожу строку к нижнему регистру, тогда не имеет значение в каком регистре ввели сложность
		switch inputLevel {
		case "easy":
			maxAttempts = 15
			maxNumber = 50
		case "medium":
			maxAttempts = 10
			maxNumber = 100
		case "hard":
			maxAttempts = 5
			maxNumber = 200
		default:
			fmt.Println("Вы ввели неверное значение - установлен уровень по умолчанию Medium")
			maxAttempts = 10
			maxNumber = 100
		}
	} else {
		fmt.Println("Вы ввели неверное значение - установлен уровень по умолчанию Medium")
		//значения устанавливаем обязательно иначе будет паника и цикл не запустится
		maxAttempts = 10
		maxNumber = 100
	}
}
