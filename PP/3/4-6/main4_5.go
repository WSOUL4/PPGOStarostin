package main

import (
	"fmt"
	"math/rand"
)

func MyArrayCut(array []int, p1, p2 int) []int {
	if p1 >= 0 {
		if p1 < len(array) {
			if p2 >= 0 {
				if p2 < len(array) {
					return array[p1:p2]
				}
			}
		}

	}
	return array
}
func MyAppend(array []int, n int) []int {
	array = append(array, n)
	return array
}
func MyDelete(array []int, n int) []int {

	if n >= 0 {
		if n < len(array) {
			array = append(array[:n], array[n+1:]...)
			return array
		}
	}
	return array
}
func MyStringCut(s string, p1, p2 int) string {
	if p1 >= 0 {
		if p1 < len(s) {
			if p2 >= 0 {
				if p2 < len(s) {
					return s[p1:p2]
				}
			}
		}

	}
	return s
}
func MyWhoIsBiggerAfterCut(s1, s2 string) string {
	if len(s1) > len(s2) {
		return "Первая больше:" + s1

	} else if len(s2) > len(s1) {
		return "Вторая больше:" + s2
	} else {
		return "Они одинаковые" + s1
	}
}
func main() {

	var numbers [5]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(1000)
	}

	fmt.Println(numbers)
	numbers2 := MyDelete(numbers[:], 2)
	fmt.Println(numbers2)
	s1 := "String 1"
	s2 := "BIGGER String the second one"
	fmt.Println(s1)
	fmt.Println(s2)
	s1 = MyStringCut(s1, 0, 6)

	fmt.Println("После обрезки:", s1)

	fmt.Println(MyWhoIsBiggerAfterCut(s1, s2))
}
