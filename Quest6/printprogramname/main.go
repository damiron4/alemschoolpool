package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0][2:]
	for _, index := range name {
		z01.PrintRune(index)
	}
	z01.PrintRune('\n')
}
