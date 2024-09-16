package main

import (
	"fmt"

	//"github.com/WSOUL4/PPpractikyGolang/tree/main/myLibs/mathutils"    // НЕ РАБОТАЕТ пишет, что нашло, но скачивать не будет, тк есть апгрейд, которого нет, у меня 1 версия, кринж, и при попытке взять ту версию, пишет, что её не существует
	"mathutils"
)

// собрать go build -o main.exe mathutils.go factcheck.go
func main() {
	var d int
	fmt.Println("Введите число:")
	/*d = 3*/ fmt.Scanf("%d\n", &d)
	d = mathutils.MyFactorial(d)
	fmt.Println("Факториал:", d)
}
