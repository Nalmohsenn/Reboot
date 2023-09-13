package piscine

func Capitalize(s string) string {
	ans := ToUpper(s)
	capital := false
	arr := []rune(ans)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			if capital == false {
				capital = true
			} else {
				arr[i] = arr[i] + 32
			}
		} else if arr[i] >= '0' && arr[i] <= '9' {
			if !capital {
				capital = true
			} else {
				continue
			}
		} else {
			capital = false
		}
	}
	return string(arr)
}
