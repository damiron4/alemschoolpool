package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ind := os.Args
	for i := 1; i < len(ind); i++ {
		ans := os.Args[i][0:]
		for _, index := range ans {
			z01.PrintRune(index)
		}
		z01.PrintRune('\n')
	}
}
