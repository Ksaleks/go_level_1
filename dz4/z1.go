package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, num)
	}
	InsertionSort(inputNums)
	fmt.Println(inputNums)
}

func InsertionSort(array []int64) {
	for i := 1; i < len(array); i++ {
		x := array[i]
		j := i
		for ; j >= 1 && array[j-1] > x; j-- {
			array[j] = array[j-1]
		}
		array[j] = x
	}
}
