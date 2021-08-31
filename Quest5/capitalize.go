package piscine

func Capitalize(s string) string {
	var up string
	str := []rune(s)
	if str[0] >= 97 && str[0] <= 122 {
		str[0] -= 32
	}
	up += string(str[0])
	for i := 1; i < len(str); i++ {
		if (str[i] >= 97 && str[i] <= 122) && !(str[i-1] >= 65 && str[i-1] <= 90 || str[i-1] >= 97 && str[i-1] <= 122 || str[i-1] >= 48 && str[i-1] <= 57) {
			str[i] -= 32
			up += string(str[i])
			continue
		} else if (str[i] >= 65 && str[i] <= 90) && (str[i-1] >= 65 && str[i-1] <= 90 || str[i-1] >= 97 && str[i-1] <= 122 || str[i-1] >= 48 && str[i-1] <= 57) {
			str[i] += 32
			up += string(str[i])
			continue
		}
		up += string(str[i])
	}
	return up
}
