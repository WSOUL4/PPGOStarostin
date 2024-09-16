package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func PrintPerson(p person) {
	fmt.Println("Имя: ", p.name, "\nВозраст: ", p.age)
}
func Birthday(p *person) {
	p.age += 1
}

// ------------------------------------------------------------------------------------------
type Shape interface {
	Area()
}

func AreaI(s Shape) { //Это то как другие Shape будут его вызывать
	s.Area() //Это функция которую надо будет переопределять для опр входимостей
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	lenght float64
}

//Методы интерфейса переопределяются немного по другому чем обычные

func (c Circle) Area() { //степень и пи есть в пакете "math", но мы же сумасшедшие тут
	var S float64
	S = c.radius * c.radius * 3.14
	fmt.Printf("Радиус окружности = %.2f\nПлощадь круга = %.2f\n", c.radius, S)
}
func (r Rectangle) Area() {
	var S float64
	S = r.lenght * r.width
	fmt.Printf("Ширина, длина = %.2f, %.2f\nПлощадь прямоугольника = %.2f\n", r.width, r.lenght, S)
}
func AreaForAll(Shapes ...Shape) {
	for _, s := range Shapes {
		s.Area()
	}
}

//------------------------------------------------------------------------------------------

type Stringer interface { //допустим, что в книге она выведет всю инфу в строковом варианте
	String()
}

func StringerPrintI(s Stringer) {
	s.String()
}

type Book struct {
	Title  string
	Author string
	Price  int
}

func (b Book) String() {

	fmt.Println("Название: ", b.Title, "\nАвтор: ", b.Author, "\nЦена: ", b.Price)
}

func main() {
	fmt.Println("ЛЮДИ")
	var alice person = person{age: 23, name: "Alice"}
	PrintPerson(alice)
	var AlPointer *person = &alice
	Birthday(AlPointer)
	PrintPerson(alice)
	fmt.Println("ФОРМЫ")
	var myCirc Circle = Circle{radius: 18.1}
	AreaI(myCirc)

	var myRec Rectangle = Rectangle{width: 11, lenght: 12.2}
	AreaI(myRec)
	fmt.Println("\nТеперь для всех")
	AreaForAll(myCirc, myRec)
	fmt.Println("КНИГИ")
	var myBook Book = Book{Title: "Некрономикон", Author: "Неизвестный", Price: 20000000}
	StringerPrintI(myBook)

}
