package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	alpha := false
	if len(s) == 0 {
		return true
	}
	for _, letter := range arr {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || (letter >= '0' && letter <= '9') {
			alpha = true
		} else {
			return false
		}
	}
	return alpha
}
