package main

import (
	"fmt"
)

func myIntSum(a, b float64) {
	fmt.Printf("Рез. суммы=%f\n", a+b)

	//return a + b
}
func myIntRed(a, b float64) {
	fmt.Printf("Рез. разн.=%f\n", a-b)

	//return a - b
}
func myIntMult(a, b float64) {
	fmt.Printf("Рез. умн.=%f\n", a*b)

	//return a * b
}
func myIntDiv(a, b float64) {
	fmt.Printf("Рез. дел.=%f\n", a/b)

	//return a / b
}

func main() {
	var a, b, d float64
	fmt.Println("Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.")
goToIn:
	fmt.Println("Введите a=")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Введите b=")
	fmt.Scanf("%f\n", &b)
	fmt.Println("1.Сумма\n2.Разность\n3.Умножение\n4.Деление")
	fmt.Scanf("%f\n", &d)
	switch d {
	case 1:
		myIntSum(a, b)
	case 2:
		myIntRed(a, b)
	case 3:
		myIntMult(a, b)
	case 4:
		myIntDiv(a, b)
	}
	goto goToIn
}
