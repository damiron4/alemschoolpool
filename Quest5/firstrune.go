package piscine

func FirstRune(s string) rune {
	ans := []rune(s)
	for _, word := range ans {
		return word
	}
	return 0
}
