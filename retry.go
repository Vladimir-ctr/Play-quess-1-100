// функция перезапуска игры
package main

import (
	"fmt"
	"strings"
)

func retry() string {

	for {
		var retryGame string
		fmt.Println("Хотите сыграть еще раз? Y - да, N - нет")
		_, err := fmt.Scan(&retryGame)

		if err != nil { // Если произошла ошибка ввода (например, пустой поток), обрабатываем
			fmt.Println("Пожалуйста введите Y или N")

			var dummy string
			fmt.Scan(&dummy) // Сброс введенных некорректных данных
			continue
		}

		retryGame = strings.ToLower(retryGame)

		if retryGame == "y" {
			levelChoice()
			game(maxAttempts, maxNumber)
			continue
		} else if retryGame == "n" {
			return "Спасибо за игру, приходите еще"
		} else {
			fmt.Println("Пожалуйста введите Y или N")
		}
	}

}
