package main

import (
	"fmt"
)

func myMean3(numbers ...int) {
	sum := 0
	leng := 0
	for _, number := range numbers {
		sum += number
		leng += 1
	}

	fmt.Printf("Среднее=%d\n", (sum / leng))

	//return a * b
}

func main() {
	fmt.Println("Написать программу, которая вычисляет среднее значение трех чисел.")
	var d1, d2, d3 int
	/*
			//var l int
			//len = 1

			//fmt.Println("Введите сколько будет чисел:")
			//fmt.Scanf("%d\n", &l)
			//const len int = l
			var numbers [len] int//длина массива должна быть константной переменной поэтому кринж по выбору кол-ва
			for i := 0; i < 3; i++{
		        fmt.Printf("Введите %d число:",i)
				fmt.Scanf("%d\n", &numbers[i])
		    }
	*/
	fmt.Println("Введите первое:")
	fmt.Scanf("%d\n", &d1)
	fmt.Println("Введите второе:")
	fmt.Scanf("%d\n", &d2)
	fmt.Println("Введите третье:")
	fmt.Scanf("%d\n", &d3)
	myMean3(d1, d2, d3)
}
