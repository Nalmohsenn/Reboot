package main 

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	arr := []rune(arguments[0])
	for i := 0; i < len(arr); i++ {
		if arr[i] != '.' && arr[i] != '/' {
			z01.PrintRune(arr[i])
		}
	}
	z01.PrintRune('\n')
}
