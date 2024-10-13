package main

import (
	"fmt"
	"strconv"
	"strings"

	//"sync"
	"time"
)

// mutex *sync.Mutex
func calc() {
	//mutex.Lock()
	var rez int
	a := <-ac
	b := <-bc
	D := <-Dc
	if D == "+" {
		rez = a + b

	} else if D == "-" {
		rez = a - b

	} else if D == "*" {
		rez = a * b

	} else if D == "/" {
		if b != 0 {
			rez = a / b

		}

	} else {
		f <- 0
	}
	f <- rez
	//mutex.Unlock()
	//timer <- true
	//return rez
}
func conv(abD []string) (int, int, string) {
	var a, b int
	var D string
	D = abD[1]
	a, _ = strconv.Atoi(abD[0])
	b, _ = strconv.Atoi(abD[2])
	return a, b, D
}
func sendto(a, b int, D string) {
	ac <- a
	bc <- b
	Dc <- D
}

var ac chan int = make(chan int)
var bc chan int = make(chan int)
var f chan int = make(chan int)
var Dc chan string = make(chan string)

func main() {
	/*defer close(ac)
	defer close(bc)
	defer close(Dc)
	defer close(f)*/

	var a, b int
	var D string
	//var mutex sync.Mutex // определяем мьютекс
	//timer := make(chan bool)
	fmt.Println("Введите простое действие...в виде (a + b)")
	//fmt.Scanf("%s", &s)
	sa := [4]string{"13 - 4", "13 + 4", "13 * 4", "13 / 4"}
	//s := "13 - 4"

	for _, el := range sa {
		abD := strings.SplitN(el, " ", 3)
		a, b, D = conv(abD)

		go sendto(a, b, D)
		time.Sleep(1)
		go calc()
		time.Sleep(1)
		fmt.Println(el, "=", <-f)
	}
	//time.Sleep(1)
	for i := 0; i < cap(ac); i++ {
		<-ac
	}
	for i := 0; i < cap(bc); i++ {
		<-bc
	}
	for i := 0; i < cap(f); i++ {
		<-f
	}
	for i := 0; i < cap(Dc); i++ {
		<-Dc
	}

	//fmt.Println(calc(a, b, D))

	//Значения передавать и получать через каналы, сама работа через треклятые рутины

	time.Sleep(1)
	//var mutex sync.Mutex // определяем мьютекс

	fmt.Print("\nend")
}
