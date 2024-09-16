package main

import (
	"fmt"
)

func myIntSum(a, b int) {
	fmt.Printf("Рез. суммы=%d\n", a+b)

	//return a + b
}
func myIntRed(a, b int) {
	fmt.Printf("Рез. разн.=%d\n", a-b)

	//return a - b
}
func myIntMult(a, b int) {
	fmt.Printf("Рез. умн.=%d\n", a*b)

	//return a * b
}
func myIntDiv(a, b int) {
	if b==0{
		fmt.Println("Алло\n")
	}else{
	fmt.Printf("Рез. дел.=%d\n", a/b)
}
	//return a / b
}

func main() {
	var a, b, d int
	fmt.Println("Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.")
goToIn:
	fmt.Println("Введите a=")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Введите b=")
	fmt.Scanf("%d\n", &b)
	fmt.Println("1.Сумма\n2.Разность\n3.Умножение\n4.Деление")
	fmt.Scanf("%d\n", &d)
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
