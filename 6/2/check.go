package main

import (
	"fmt"
)

func fibonacciIterative2(n int) int {
	if n <= 1 {
		return n
	}
	var n2, n1 = 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	//Channel666 <- n1
	//close(Channel666)
	return n1
}
func fibonacciIterative(Channel666 chan int, n int) int {
	if n <= 1 {
		return n
	}
	var n2, n1 = 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	Channel666 <- n1
	close(Channel666)
	return n1
}

/*
	func g10f() {
		//n:=10
		//var array [10]int

		for I := 0; I < 10; I++ {
			//array[I] = fibonacciIterative(I + 1)
			Channel666 <- fibonacciIterative(I + 1)
			close(Channel666)
		}
		//Отправить массив теперь в канал

}
*/
func greadchan(Channel666 chan int) {
	fmt.Printf("%d\n", <-Channel666)

}

//var Channel666 chan int

func main() {
	for i := 1; i <= 10; i++ {

		fmt.Printf("\n%d\n", fibonacciIterative2(i))
		//fmt.Print('\n')

	}

	fmt.Print("\nend")
}
