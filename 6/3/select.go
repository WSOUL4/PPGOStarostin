package main

import (
	"fmt"
	"math/rand"
	"time"
)

func oddeven() {
	defer close(r)
	for i := 0; i < 5; i++ {
		//d := <-c
		select {
		case d := <-c:
			//d:= <-c
			if d%2 == 0 {
				//fmt.Println(d, "Ч")
				r <- "Чётное"
			} else if d%2 > 1 {
				//fmt.Println(d, "О")
				r <- "Ноль"
			} else if d%2 == 1 {
				//fmt.Println(d, "Н")
				r <- "Не чётное"
			}

		}
		time.Sleep(1 * time.Second)
	}
}

func getrand() {
	defer close(c)
	for i := 1; i < 5; i++ {
		ran := rand.Intn(50)
		select {
		case c <- ran:
			//fmt.Printf("Загрузили рандом %d\n", ran)
			c <- ran
			time.Sleep(1 * time.Second)
			//fmt.Printf("Загрузили рандом check %d\n", <-c)
		}
	}

}
func cover() {

	//time.Sleep(1 * time.Second)
	go getrand()
	//
	//time.Sleep(1 * time.Second)
	go oddeven()

	//time.Sleep(1 * time.Second)
	/*
		for i := 0; i < 10; i++ {
			select {
			case d := <-c:
				//fmt.Println("Case print")
				fmt.Println(d)
			case s := <-r:
				//fmt.Println("Case print")
				fmt.Println(s)
				//default:
				//fmt.Println("default")

			}
		}
	*/

}

var c chan int = make(chan int, 5)
var r chan string = make(chan string, 5)

func main() {

	go cover()
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case d := <-c:
			//fmt.Println("Case print")
			fmt.Println(d)
		case s := <-r:
			//fmt.Println("Case print")
			fmt.Println(s)
			//default:
			//fmt.Println("default")

		}
	}
	//close(Channel666)
	//close(Channel2)
	fmt.Print("\nend\n")
}
