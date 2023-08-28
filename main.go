package main

import "fmt"

const (
	balance    string = "Скобки сбалансированы"
	notBalance string = "Скобки не сбалансированы"
)

func main() {
	var str string
	fmt.Scanln(&str)

	// Проверка на пустую строку. При отсутствий скобок считается, что строка сбалансирована.
	if str == "" {
		fmt.Println(balance)
		return
	}

	// Проверка на четность скобок. Если количество скобок нечетное, то строка считается не сбалансированной.
	if len(str)%2 != 0 {
		fmt.Println(notBalance)
		return
	}

	// Стек левых скобок.
	leftBrackets := make([]byte, 0, len(str)/2)

	for i := range str {
		if str[i] == '{' || str[i] == '(' || str[i] == '[' {
			// Запись левых скобок в стек.
			leftBrackets = append(leftBrackets, str[i])
			continue
		}

		if str[i] == '}' || str[i] == ')' || str[i] == ']' {
			// Проверка на наличие левых скобок в сетке.
			if len(leftBrackets) > 0 {
				// Получение индекса верхнего значения стека для левых скобок.
				n := len(leftBrackets) - 1

				// Определение совпадает ли тип правой и левой скобки.
				// Работа с байтами (ASCII символы):
				// Двойка вычитается для сравнение фигурных ("{" - 123, "}" - 125) и квадратных скобок ("[" - 91, "]" - 93)
				// Единица вычитается для сравнения круглых ("(" - 40, ")" - 41) скобок
				if leftBrackets[n] == str[i]-2 || leftBrackets[n] == str[i]-1 {
					// Удаление правой скобки из вершины стека.
					leftBrackets = leftBrackets[:n]
				} else {
					fmt.Println(notBalance)
					return
				}
			} else {
				fmt.Println(notBalance)
				return
			}
		}
	}

	// Проверка на пустоту стека левых скобок.
	// Если левые скобки остались, значит, строка считается не сбалансированной. Левых скобок больше правых.
	if len(leftBrackets) != 0 {
		fmt.Println(notBalance)
		return
	}

	fmt.Println(balance)
}
