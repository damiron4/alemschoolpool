package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var temp int = n
	var count int = 1
	if n < 0 {
		n *= -1
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for i := 0; temp > 0; i++ {
		if temp == 0 {
			break
		} else if temp <= 9 {
			break
		}
		temp /= 10
		count *= 10
	}
	for ok := true; ok; ok = (n > 0) {
		if n == 0 {
			break
		} else if n <= 9 {
			pop := n
			z01.PrintRune(rune(pop))
		}
		pop := n / count
		n %= count
		count /= 10
		z01.PrintRune(rune(pop))
	}
}
