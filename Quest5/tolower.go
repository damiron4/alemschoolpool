package piscine

func ToLower(s string) string {
	slice := []rune(s)
	for i := 0; i < len(s); i++ {
		if IsUpper(string(slice[i])) {
			slice[i] = rune(slice[i] + 32)
		}
	}
	return string(slice)
}
