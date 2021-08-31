package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if i < j && j < k {
					if i == '7' && j == '8' && k == '9' {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						z01.PrintRune(rune('\n'))
						break
					}
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(','))
					z01.PrintRune(rune(' '))
				}
			}
		}
	}
}
