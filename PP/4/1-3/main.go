package main

import (
	"fmt"
)

func PrintMap(people map[string]int) {
	for key, value := range people {
		fmt.Println(key, value)
	}
}

func AskToAdd(people map[string]int) map[string]int {
	fmt.Println("Введите имя человека: ")
	var s string
	fmt.Scanf("%s\n", &s)
	fmt.Println("Введите возраст его/её: ")
	var i int
	fmt.Scanf("%d\n", &i)
	people[s] = i
	return people
}
func MeanPeople(people map[string]int) int {
	sum := 0
	leng := 0
	for key, value := range people {
		sum += value
		leng += 1
		key = key //он меня бесит можно _
	}
	return sum / leng

}
func KillPeople(people map[string]int, name string) map[string]int {
	delete(people, name)
	return people
}
func main() {

	var people = map[string]int{"Tom": 10, "Bob": 12}
	people = AskToAdd(people)
	PrintMap(people)
	fmt.Println("Cредний возраст:", MeanPeople(people))
	people = KillPeople(people, "Tom")
	PrintMap(people)
}
