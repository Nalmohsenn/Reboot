package piscine

func IsLower(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		if arr[i] < 'a' || arr[i] > 'z' {
			return false
		}
	}
	return true
}
