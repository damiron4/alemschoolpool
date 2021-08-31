package main

import "github.com/01-edu/z01"

func main() {
	for i := 122; i >= 97; i-- {
		character := rune(i)
		z01.PrintRune(character)
	}
	z01.PrintRune('\n')
}
