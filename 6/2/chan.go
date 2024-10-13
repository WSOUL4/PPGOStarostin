package main

import (
	"fmt"
	"time"
)

func fibonacciIterative(n int) {
	//defer close(Channel666) // закрываем канал после завершения горутины
	if n <= 1 {
		//return n
	}
	var n2, n1 = 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}

	Channel666 <- n1
	//close(Channel666)
	//return n1
}

func greadchan() {
	fmt.Println(<-Channel666)
	//Channel666 = make(chan int)

}

var Channel666 chan int = make(chan int, 10)

func sender() {
	for i := 1; i <= 10; i++ {
		fibonacciIterative(i)
	}
	close(Channel666)
}
func receiver() {
	for i := 1; i <= 10; i++ {
		greadchan()
	}
}
func main() {
	//Channel666 := make(chan int)

	go sender()

	go receiver()

	//<-Channel666 // ожидаем закрытия канала

	time.Sleep(10)
	fmt.Print("\nend\n")
	for i := 0; i < cap(Channel666); i++ {
		if val, opened := <-Channel666; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
	fmt.Print("\nend")
}
