package main

import (
	"fmt"
)

func myMean3(numbers ...int) int {
	sum := 0
	leng := 0
	for _, number := range numbers {
		sum += number
		leng += 1
	}

	//fmt.Printf("Среднее=%d\n", (sum / leng))

	return (sum / leng)
}

func main() {
	var d1, d2 int

	fmt.Println("Введите первое:")
	fmt.Scanf("%d\n", &d1)
	fmt.Println("Введите второе:")
	fmt.Scanf("%d\n", &d2)

	fmt.Printf("Среднее: %d\n", myMean3(d1, d2))
}
