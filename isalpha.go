package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= '0' && arr[i] <= '9') {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
