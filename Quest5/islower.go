package piscine

func IsLower(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			count++
		}
	}
	return count == len(s)
}
