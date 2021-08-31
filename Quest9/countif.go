package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, ans := range tab {
		if f(ans) {
			count++
		}
	}
	return count
}
