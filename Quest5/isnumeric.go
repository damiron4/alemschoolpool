package piscine

func IsNumeric(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			count++
		}
	}
	return count == len(s)
}
