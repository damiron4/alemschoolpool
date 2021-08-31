package piscine

func Atoi(s string) int {
	var b int = 0
	c := 0
	a := 0
	d := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	for i := 0; i < a; i++ {
		if s[i] == '-' {
			if i > 0 {
				return 0
			}
			c = 1
			d++
			continue
		}
		if s[i] == '+' {
			if i > 0 {
				return 0
			}
			d++
			continue
		}
		if s[i] == ' ' {
			return 0
		}
		b = b*10 + int(s[i]) - '0'
		if int(s[i])-'0' > 9 || int(s[i])-'0' < 0 {
			return 0
		}
	}
	if d > 1 {
		return 0
	}
	if c == 1 {
		b *= -1
	}
	return b
}
