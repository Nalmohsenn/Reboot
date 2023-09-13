package piscine

func TrimAtoi(s string) int {
	arr := []rune(s)
	ans := 0
	min := false
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			ans = (ans*10 + int(rune(arr[i]-'0')))
		}
		if arr[i] == '-' && ans == 0 {
			min = true
		}
	}
	if min == true {
		ans = ans * -1
		return ans
	} else {
		return ans
	}
}
