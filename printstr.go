package piscine

import (
	"github.com/01-edu/z01.PrintRune"
)

func PrintStr(s string) {
	ar := []rune(s)
	for _, t := range ar {
		z01.PrintRune(t)
	}
}
