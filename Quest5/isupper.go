package piscine

func IsUpper(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			count++
		}
	}
	return count == len(s)
}
