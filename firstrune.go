package piscine

func FirstRune(s string) rune {
	array := []rune(s)
	return array[len(s)-1]
}
