package main

import (
	"fmt"
	"strings"
)

func main() { //можно импортировать os и там os.SdtIn, что бы пробелы считались вместо параметра в скане
	var s string
	fmt.Println("Введите строку: ")
	fmt.Scanf("%s", &s)
	fmt.Println("Изменённая строка: ", strings.ToUpper(s))
}
