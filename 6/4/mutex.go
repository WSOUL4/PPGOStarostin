package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 0

func incr1(i int, mutex *sync.Mutex, timer chan bool) {

	mutex.Lock()
	count += 1
	fmt.Println("Goroutine", "1#", i, "=", count)
	mutex.Unlock()
	timer <- true

}
func incr2(i int, mutex *sync.Mutex, timer chan bool) {

	mutex.Lock()
	count += 100
	fmt.Println("Goroutine", "2#", i, "=", count)
	mutex.Unlock()
	timer <- true

}

var ch bool = false

func main() {
	timer := make(chan bool) // канал
	//ch <- false

	var mutex sync.Mutex // определяем мьютекс
	//mutex.Unlock()
	//fmt.Println(mutex.TryLock())
	for i := 0; i < 5; i++ {
		go incr1(i, &mutex, timer)
		time.Sleep(1)
		go incr2(i, &mutex, timer)
		time.Sleep(1)

	}
	for i := 0; i < 10; i++ {
		<-timer
	}

	fmt.Println("end")
}
