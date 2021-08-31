package piscine

func Compare(a, b string) int {
	slice1 := []byte(a)
	slice2 := []byte(b)
	x := len(a)
	y := len(b)
	if x >= y {
		for i := 0; i < x; i++ {
			if slice1[i] > slice2[i] { // str1 > str2
				return 1
			} else if slice1[i] == slice2[i] { // str1 == str2
				if i+1 > y-1 && i+1 < x-1 {
					return 1
				}
				continue
			} else { // str1 < str2
				return -1
			}
		}
		return 0
	} else {
		for i := 0; i < y; i++ {
			if slice1[i] > slice2[i] { // str1 > str2
				return 1
			} else if slice1[i] == slice2[i] { // str1 == str2
				if i+1 > x-1 && i+1 < y-1 { // x = 6, y = 6, i = 5,
					return 1
				}
				continue
			} else { // str1 < str2
				return -1
			}
		}
		return 0
	}
}
