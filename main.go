package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error 0")
	}
	return a / b, nil
}

func main() {
	for {
		var num1, num2 int
		var otvet string
		_, err := fmt.Scan(&num1, &num2)
		if err != nil {
			fmt.Println("Error, text")
			break
		}
		rezult, err := divide(num1, num2)
		if err != nil {
			fmt.Println("Error")
			break
		}
		switch {
		case rezult > 10:
			fmt.Println("rezult > 10")

		case rezult <= 1 && rezult >= 10:
			fmt.Println("0 < rezult < 10")

		default:
			fmt.Println("rezult < 0 or smoll")
		}
		fmt.Println("Повторить операцию? YES or NO")
		fmt.Scan(&otvet)
		if otvet == "YES" {
			fmt.Println("YES")
		} else if otvet == "NO" {
			fmt.Println("NO")
			break
		} else {
			continue
		}
	}
}
