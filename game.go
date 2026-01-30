// Выносим цикл в отдельную функцию для переиспользования
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func game(maxAttempts int, maxNumber int) (bool, int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(maxNumber) + 1

	var inputNumber int
	var printInputNumber []int // создаем переменную являющуюся срезом P.S если инициализировать слайс через make то мы уже к имеющимся значениям будем добавлять числа
	var usedAttempts int

	fmt.Printf("Я загадал число от 1 до %d. У тебя %d попыток чтобы отгадать! \n", maxNumber, maxAttempts)
	fmt.Println("Введите число")

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Попытка %d: введите число: ", attempt)
		// Проверка ввода
		_, err := fmt.Scan(&inputNumber) // сохраняем введеное число в переменную inputNumber. _, err - способ получить два числа, но число нам не нужно, поэтому ставим заглушку
		if err != nil {                  // Если err не равна nil — значит, пользователь ввёл не число (например, буквы или символы). Тогда нужно обработать эту ситуацию
			fmt.Println("Пожалуйста введите корректное число")
			fmt.Scanf("%s") // Сброс введенных некорректных данных
			attempt--
			continue

		}

		usedAttempts = attempt                                   // запоминаем колличество реально использованных попыток
		printInputNumber = append(printInputNumber, inputNumber) //принимаем слайс и записываем в него значение

		if hints(inputNumber, randomNumber) {
			return true, usedAttempts // return проверяет какое булевое значение вернула функция hints и если true тогда выходим из функции game // после добавления true если угадал возвращаем
			// после добавления true если угадал возвращаем побуду и колличество попыток
		}

		fmt.Printf("Числа которые вы ввели ранее: %d \n", printInputNumber) // Выводим список чисел

	}

	fmt.Println("Попытки кончились, вы проиграли")
	fmt.Printf("Загаданное число было: %d\n", randomNumber)
	return false, usedAttempts

}
