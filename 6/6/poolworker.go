package main

import (
	"fmt"
	"time"
)

type Work interface {

	// Note that we will
	// declare the methods to be used
	// later here in this
	// interface
	Read()
	Reverse()
}
type Worker struct {
	id int
}

func (x Worker) Reverse() {
	// this method will access the variable
	// food in Animal class

	go MyReverse(x)
}
func (x Worker) Read(s string) {
	// this method will access the variable
	// food in Animal class

	Origs <- s
	fmt.Println(x.id, "#Прочёл-", s)
}

func MyReverse(x Worker) {
	s := <-Origs
	slen := len(s)
	runes := []rune(s)
	for i, c := range s {
		// действия
		newi := slen - 1 - i
		runes[newi] = c
	}
	fmt.Println(x.id, "#Повернул-", string(runes))
	Revs <- string(runes)

}

var Revs chan string = make(chan string)
var Origs chan string = make(chan string)

func main() {

	pool := [3]Worker{Worker{id: 1}, Worker{id: 2}, Worker{id: 3}}
	for iter, man := range pool {
		if iter == 0 {
			go man.Read("Poolakoka")
			time.Sleep(1)
		} else if iter == 1 {
			go man.Reverse()
			time.Sleep(1)
		} else if iter == 2 {
			go man.Read("34545454")
			time.Sleep(1)
			go man.Reverse()
			time.Sleep(1)
		}
	}
	for i := 0; i < cap(Revs); i++ {
		<-Revs
	}
	for i := 0; i < cap(Origs); i++ {
		<-Origs
	}
	fmt.Print("\nend")
}
