package piscine

func TrimAtoi(s string) int {
	var b int = 0
	c := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			if c != 0 {
				continue
			}
			if count > 0 {
				continue
			}
			count++
		}
		if int(s[i])-'0' > 9 || int(s[i])-'0' < 0 {
			continue
		}
		b = b*10 + int(s[i]) - '0'
		if count > 0 {
			continue
		}
		c++
	}
	if count == 1 {
		b *= -1
	}
	return b
}
