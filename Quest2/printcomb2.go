package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for x := '0'; x <= '9'; x++ {
				for y := '0'; y <= '9'; y++ {
					if i <= x && j <= y || (i < x && j > y) {
						if i == '9' && j == '8' && x == '9' && y == '9' {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(' '))
							z01.PrintRune(rune(x))
							z01.PrintRune(rune(y))
							z01.PrintRune(rune('\n'))
							break
						} else if i == x && j == y {
							continue
						} else {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(' '))
							z01.PrintRune(rune(x))
							z01.PrintRune(rune(y))
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}
					}
				}
			}
		}
	}
}
