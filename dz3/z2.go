package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Println("Введите число")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка :", err.Error())
		os.Exit(1)
	}
	for i := 2; i <= n; i++ {
		cnt := 0
		del := 2
		for del < i {
			if i%del == 0 {
				cnt++
			}
			del++
		}
		if cnt == 0 {
			fmt.Printf("%v - простое число \n", i)
		}
	}
}
