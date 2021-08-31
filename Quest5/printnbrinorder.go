package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr [10]int
	b := n
	if n == 0 {
		z01.PrintRune(48)
	}
	for b > 0 {
		arr[b%10]++
		b /= 10
	}
	for c := 0; c <= 9; c++ {
		for i := 0; i < arr[c]; i++ {
			z01.PrintRune(rune(c + 48))
		}
	}
}
