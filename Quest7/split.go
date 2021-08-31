package piscine

func Split(s, sep string) []string {
	var str []string
	temp := []rune(s)
	num_sp := 0
	ind := 0
	smth := 0
	for i := 0; i < len(s); i++ {
		count := 0
		if temp[i] == rune(sep[0]) {
			for j, k := 0, i; j < len(sep); j, k = j+1, k+1 {
				if temp[k] == rune(sep[j]) {
					count++
				}
			}
			if count == len(sep) {
				num_sp++
			}
		}
	}
	for i := 0; i < len(s); i++ {
		count := 0
		if temp[i] == rune(sep[0]) {
			for k, j := i, 0; j < len(sep); k, j = k+1, j+1 {
				if temp[k] == rune(sep[j]) {
					count++
				}
			}
			if count == len(sep) {
				ind = smth
				if ind != 0 {
					ind += len(sep)
				}
				smth = i
				str = append(str, string(temp[ind:smth]))
				num_sp--
			}
		}
		if num_sp == 0 {
			str = append(str, string(temp[smth+len(sep):]))
			break
		}
	}
	return str
}
