package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	slice := []byte(s)
	for index := range s {
		z01.PrintRune(rune(slice[index]))
	}
}
