package piscine

func IsPrintable(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] > 0 && s[i] <= 31 {
			continue
		}
		count++
	}
	return count == len(s)
}
