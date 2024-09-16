package main

import (
	"fmt"
)

func mySum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number

	}
	return sum
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func main() {
	fmt.Println("Сколько будет чисел: ")
	var n int
	if m, err := Scan(&n); m != 1 { //Можно доп. условия через ;
		panic(err)
	}
	fmt.Println("Введите числа через пробел: ")
	all := make([]int, n)
	ReadN(all, 0, n)

	fmt.Println(all)
	fmt.Println("Сумма: ", mySum(all))
}
