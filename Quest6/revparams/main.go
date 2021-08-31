package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ind := os.Args
	for i := len(ind) - 1; i > 0; i-- {
		ans := os.Args[i][0:]
		for _, index := range ans {
			z01.PrintRune(index)
		}
		z01.PrintRune('\n')
	}
}
