package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		var number1, number2 int
		var question string
		_, err := fmt.Scan(&number1, &number2)
		if err != nil {
			fmt.Println("Error")
			continue
		}

		rezult, err := divide(number1, number2)
		if err != nil {
			fmt.Println("Error")
			continue
		}

		switch {
		case rezult > 10:
			fmt.Println("Результат большой")

		case rezult <= 10 && rezult >= 1:
			fmt.Println("Результат средний")

		default:
			fmt.Println("Результат маленький или ноль")
		}
		fmt.Scan(&question)
		if question == "yes" {
			fmt.Println("Заново")
		} else if question == "no" {
			break
		} else {
			continue
		}
	}

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error")
	}
	return a / b, nil
}
