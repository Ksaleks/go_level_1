package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var op string
	var a, b float64
	fmt.Println("Введите два числа и оператор между ними")
	_, err := fmt.Scan(&a, &op, &b)
	if err != nil {
		fmt.Println("Числа дожны быть типа float оператор типа string")
		os.Exit(1)
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "^":
		fmt.Println(math.Pow(a, b))
	default:
		fmt.Println("Введен неверный оператор. Поддерживается +, -, *, /, ^")
	}
}
