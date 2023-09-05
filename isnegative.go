package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb < 0 {
		z01.printRune('T')
		z01.printRune('\n')
	} else {
		z01.printRune('F')
		z01.printRune('\n')
	}
}
