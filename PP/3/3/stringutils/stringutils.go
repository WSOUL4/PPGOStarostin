package stringutils

func MyReverse(s string) string {
	slen := len(s)
	runes := []rune(s)
	for i, c := range s {
		// действия
		newi := slen - 1 - i
		runes[newi] = c
	}
	return string(runes)
}
