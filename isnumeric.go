package piscine

func IsNumeric(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			count++
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
