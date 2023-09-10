package piscine

func FirstRune(s string) rune {
	array := []Rune(s)
	return array[len(s)-1]
}
