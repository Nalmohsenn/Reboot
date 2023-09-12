package piscine

func ToUpper(s string) string {
	slices := []rune(s)
	for i := 0; i < len(s); i++ {
		if slices[i] >= 97 && slices[i] <= 122 {
			slices[i] = rune(slices[i] - 32)
		}
	}
	return string(slices)
}
