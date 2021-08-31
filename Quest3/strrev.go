package piscine

func StrRev(s string) string {
	a := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	slice := []byte(s)
	for i, b := a-1, 0; i >= 0; i, b = i-1, b+1 {
		slice[b] = s[i]
	}
	str := string(slice)
	return str
}
