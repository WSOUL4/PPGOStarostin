package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func RecArea(r Rectangle) float64 {
	if r.height*r.width < 0 {
		return r.height * r.width * -1
	}
	return r.height * r.width
}
func main() {
	var w, h float64
	fmt.Println("Введите длинну прямоугольника:")
	fmt.Scanf("%s\n", &w)
	fmt.Println("Введите ширину прямоугольника:")
	fmt.Scanf("%s\n", &h)
	var r Rectangle = Rectangle{w, h}
	fmt.Println(RecArea(r))
}
