package main

//Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".
import (
	"fmt"
)

func mySign(d int) string {
	if d == 0 {
		return "Zero"
	} else if d > 0 {
		return "Positive"
	} else if d < 0 {
		return "Negative"
	}
	return "Wrong stuff"
}
func main() {
	var d int
	fmt.Println("Введите число:")
	fmt.Scanf("%d\n", &d)
	fmt.Println(mySign(d))
	//fmt.Println("")
}
