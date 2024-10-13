package mathutils

//Имена функци доллжны начинаться с большой буквы или их не видно другому пакету

func MyFactorial(d int) int {
	if d > 0 {
		a := 1
		for i := 2; i <= d; i++ {
			a = a * i
		}
		return a
	}
	return -1
}
