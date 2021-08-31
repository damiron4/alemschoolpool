package piscine

func StrLen(s string) int {
	a := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	return a
}
