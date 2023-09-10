package piscine

func FirstRune(s string) rune {
	arr := []rune(s)
	return arr[len(s)-1]
}
