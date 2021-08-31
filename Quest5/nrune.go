package piscine

func NRune(s string, n int) rune {
	ans := []rune(s)
	if n > len(s) || n-1 < 0 {
		return 0
	}
	return ans[n-1]
}
