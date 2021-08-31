package piscine

func BasicAtoi2(s string) int {
	var b int = 0
	a := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	for i := 0; i < a; i++ {
		if s[i] == ' ' {
			return 0
		}
		b = b*10 + int(s[i]) - '0'
		if int(s[i])-'0' > 9 || int(s[i])-'0' < 0 {
			return 0
		}
	}
	return b
}
