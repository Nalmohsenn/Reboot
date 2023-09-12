package piscine

func IsLower(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		if arr[i] >= 'A' || arr[i] <= 'Z' {
			return false
		}
	}
	return true
}
