package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("Введите да что угодно:")
	fmt.Scanf("%s\n", &s)
	fmt.Println(len(s))
}
