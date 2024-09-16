package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Написать программу, которая выводит текущее время и дату.")
	//лэауты сломаны 2-день    1-месяц 2006-год 3-часы am  4-минуты  5-секунды без нуля  15- часы 24 часовой формат АХАЪХАХАХАХАХ
	layout := "02-01-2006 15:04:05"
	layout_date := "02/01/2006"
	layout_time := "15:04:05"
	today := time.Now()
	//yr, mon, day := today.Date()
	//hr, min, sec := today.Clock()
	fmt.Println(today.String())
	fmt.Println(today.Format(layout))
	fmt.Println("Дата:", today.Format(layout_date), " \t Время: ", today.Format(layout_time))
	fmt.Println("Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.")
	var myInt int
	var myFloat float64
	var myString string
	var myFlag bool
	myInt = 13
	myFloat = 13.13
	myString = "13,13"
	myFlag = true
	fmt.Printf("Целое: %d\n", myInt)
	fmt.Printf("Десятичное: %.3f\n", myFloat) //есть %g, %f для других экспонент , точность через приставку с точкой и числом
	fmt.Printf("Строка: %s\n", myString)
	fmt.Printf("Флаг: %t\n", myFlag)

	fmt.Println("Использовать краткую форму объявления переменных для создания и вывода переменных.")
	//уже применял
	myShort := 12
	fmt.Printf("Краткая:%d", myShort)

}
