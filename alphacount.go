package piscine

import (
	"github.com/01-edu/z01"
)

func AlphaCount(s string) int {
	count := 0
	arr := []rune(s)

	for i := 0; i < len(s); i++ {
		if (arr[i] >= 'A' && arr <= 'Z') || (arr[i] >= 'a' && arr <= 'z') {
			count++
		}
	}
	return count
}
