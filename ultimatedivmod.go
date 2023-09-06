package piscine

func UltimateDivMod(a *int, b *int) {
	n1 := *a / *b
	n2 := *a % *b
	*a = n1
	*b = n2
}
