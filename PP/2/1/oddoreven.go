package main

//Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
import (
	"fmt"
)

func main() {
	var d int
	fmt.Println("Введите число:")
	fmt.Scanf("%d\n", &d)
	if d%2 == 0 {
		fmt.Println("Чётное")
	} else if d == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Не чётное")
	}
	//fmt.Println("")
}
