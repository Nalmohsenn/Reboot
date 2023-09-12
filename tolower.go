package piscine

func ToLower(s string) string {
	slices := []byte(s)
	for i := 0; i < len(s); i++ {
		if slices[i] >= 65 && slices[i] <= 90 {
			slices[i] = byte(slices[i] + 32)
		}
	}
	return string(slices)
}
