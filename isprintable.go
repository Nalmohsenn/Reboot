package piscine

func IsPrintable(s string) bool {
	count := 0
	array := []rune(s)
	for i := 0; i < len(s); i++ {
		if array[i] >= 32 && array[i] <= 127 {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
