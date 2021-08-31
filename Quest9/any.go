package piscine

func Any(f func(string) bool, a []string) bool {
	for _, ans := range a {
		if f(ans) {
			return true
		}
	}
	return false
}
