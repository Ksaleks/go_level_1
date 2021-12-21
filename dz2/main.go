package main

import (
	"fmt"
	"math"
)

func main() {
	var number int8
	fmt.Println("Выбирите тип расчета и введите его номер")
	fmt.Println("1. Вычисление площади прямоугольника")
	fmt.Println("2. Вычисление диаметра и длины окружности по заданной площади")
	fmt.Println("3. Вывод цифр, соответствующих количеству сотен, десятков и единиц трехзначного числа")
	fmt.Scan(&number)

	switch number {
	case 1:
		square()
	case 2:
		circleDiameterLength()
	case 3:
		splittingNumber()
	default:
		fmt.Println("Нет такого варианта расчета")
	}
}

func square() {
	var a, b, s int
	fmt.Println("Введите стороны прямоугольника {a b}")
	fmt.Scan(&a, &b)
	s = a * b
	fmt.Println("Площадь прямоугольника равна ", s)
}

func circleDiameterLength() {
	var s, d, l float64
	fmt.Println("Введите площадь окружности")
	fmt.Scan(&s)
	d = 2 * math.Sqrt(s/math.Pi)
	l = 2 * math.Sqrt(s*math.Pi)
	fmt.Println("Диаметр окружности: ", d)
	fmt.Println("Длина окружности: ", l)
}

func splittingNumber() {
	var hundreds, dozens, units, number int
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)
	units = number % 10
	number = (number - units) / 10
	dozens = number % 10
	number = (number - dozens) / 10
	hundreds = number
	fmt.Println("Сотни: ", hundreds)
	fmt.Println("Десятки: ", dozens)
	fmt.Println("Единицы: ", units)
}
