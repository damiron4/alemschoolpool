package piscine

func SplitWhiteSpaces(s string) []string {
	var str []string
	temp := []rune(s)
	count := 0
	ind := 0
	num_sp := 0
	for i := 0; i < len(s); i++ {
		if temp[i] == ' ' {
			num_sp++
		}
	}
	for j := 0; j < len(s); j++ {
		if temp[j] == ' ' {
			ind = count
			if count != 0 {
				ind += 1
			}
			count = j
			if string(temp[ind:count]) != "" {
				str = append(str, string(temp[ind:count]))
			}
			num_sp--
		}
		if num_sp == 0 {
			str = append(str, string(temp[count+1:]))
			break
		}
	}
	return str
}
