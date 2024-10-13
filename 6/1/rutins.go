package main

import (
	"fmt"
	"math/rand"

	"time"
)

var gsum = 0
var RandI = 0
var gf = 0

func main() {
	go Rand()
	time.Sleep(20)

	go mySum(10, RandI)
	go MyFactorial(RandI)

	//fmt.Print(mathutils.MyFactorial(1))
	time.Sleep(2)
	fmt.Println("The End")
}
func MyFactorial(d int) int {
	//d = RandI
	if d > 0 {
		a := 1
		for i := 2; i <= d; i++ {
			a = a * i
		}
		gf = a
		fmt.Printf("f%d\n", gf)
		return gf
	}
	return -1
}
func mySum(numbers ...int) int {
	sum := 0
	//leng := 0
	for _, number := range numbers {
		sum += number
		//leng += 1
	}
	gsum = gsum + sum
	fmt.Printf("s%d\n", gsum)
	return sum
}
func Rand() int {
	RandI = rand.Intn(10)
	fmt.Printf("r%d \n", RandI)
	return RandI
}
