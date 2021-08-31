package piscine

func LastRune(s string) rune {
	a := 0
	ans := []rune(s)
	for _, word := range ans {
		if word != 0 {
			a++
		}
	}
	return ans[a-1]
}
