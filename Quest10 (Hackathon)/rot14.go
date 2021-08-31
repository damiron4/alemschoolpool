package piscine

func Rot14(s string) string {
	var str string
	for _, ans := range s {
		if ans >= 97 && ans <= 122 {
			if ans >= 'm' {
				ans -= 12
			} else {
				ans += 14
			}
		} else if ans >= 65 && ans <= 90 {
			if ans >= 'M' {
				ans -= 12
			} else {
				ans += 14
			}
		}
		str += string(ans)
	}
	return str
}
