package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str [8]string
	ind := os.Args
	for i, j := 1, 0; i < len(ind); i, j = i+1, j+1 {
		str[j] += ind[i]
	}
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if str[j] > str[j+1] {
				temp := str[j]
				str[j] = str[j+1]
				str[j+1] = temp
			}
		}
	}
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			if str[i][j] == 1 {
				continue
			}
			z01.PrintRune(rune(str[i][j]))
		}
		z01.PrintRune('\n')
	}
}
