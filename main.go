package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var maxAttempts int
var maxNumber int

func main() {
	fmt.Println("Добро пожаловать в игру угадай число! Пожалуйста введите уровень сложности: Easy, Medium или Hard")

	levelChoice()

	game(maxAttempts, maxNumber)

	retry()

}

// Выносим цикл в отдельную функцию для переиспользования
func game(maxAttempts int, maxNumber int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(maxNumber) + 1

	var inputNumber int
	var printInputNumber []int // создаем переменную являющуюся срезом P.S если инициализировать слайс через make то мы уже к имеющимся значениям будем добавлять числа

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

		printInputNumber = append(printInputNumber, inputNumber) //принимаем слайс и записываем в него значение

		if hints(inputNumber, randomNumber) {
			return
		}

		fmt.Printf("Числа которые вы ввели ранее: %d \n", printInputNumber) // Выводим список чисел

	}

	fmt.Println("Попытки кончились, вы проиграли")
	fmt.Printf("Загаданное число было: %d\n", randomNumber)
}

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

// функция выбора уровня сложности
func levelChoice() {

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

// пишем дополнительную функцию abs для расчета разницы по модулю (т.е функция возвращает число всегда положительное)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// подсказки
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
