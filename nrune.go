package piscine

func NRune(s string, n int) rune {
	length := len(s)
	if n > 0 && n <= length {
		arr := []rune(s)
		return arr[n-1]
	} else {
		return 0
	}
}
