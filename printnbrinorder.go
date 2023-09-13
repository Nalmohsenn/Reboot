package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	nb := n
	count := true
	arr := [19]int{}
	countnb := 0
	if nb == 0 {
		z01.PrintRune(rune(nb + 48))
	} else {
		for i := 0; count != false; i++ {
			if nb != 0 {
				arr[i] = nb % 10
				nb = nb / 10
				countnb++
			} else {
				countnb = false
			}
		}
		temp := 0
		for j := 0; j < countnb; j++ {
			for v := j + 1; v < countnb; v++ {
				if arr[j] > arr[v] {
					temp = arr[j]
					arr[j] = arr[v]
					arr[v] = temp
				}
			}
		}
		for _, t := range arr {
			if countnb > 0 {
				z01.PrintRune(rune(t + 48))
				countnb--
			}
		}
	}
}
