package piscine

func AlphaCount(s string) int {
	count := 0
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		if (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= 'a' && arr[i] <= 'z') {
			count++
		}
	}
	return count
}
