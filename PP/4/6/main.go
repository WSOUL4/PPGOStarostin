package main

import (
	"fmt"
)

func myRev(numbers []int, n int) []int {
	d := make([]int, n)

	for i, v := range numbers {
		// действия
		i2 := n - 1 - i
		d[i2] = v
	}
	return d
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
	all := make([]int, n) //Прекрасная команда для создания массивов не используя константные значения
	ReadN(all, 0, n)

	fmt.Println(all)
	fmt.Println("Развёрнутый: ", myRev(all, n))
}
