package piscine

func BasicAtoi(s string) int {
	var b int = 0
	a := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	for i := 0; i < a; i++ {
		b = b*10 + int(s[i]) - '0'
	}
	return b
}
