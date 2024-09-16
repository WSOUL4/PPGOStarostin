package main

import (
	"fmt"

	"stringutils"
)

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scanf("%s\n", &s)
	fmt.Println(stringutils.MyReverse(s))

}
