package main

import "os"

func Atoi(s string) int {
	var b int = 0
	c := 0
	a := 0
	d := 0
	for _, index := range s {
		if index != 0 {
			a++
		}
	}
	for i := 0; i < a; i++ {
		if s[i] == '-' {
			if i > 0 {
				return 0
			}
			c = 1
			d++
			continue
		}
		if s[i] == '+' {
			if i > 0 {
				return 0
			}
			d++
			continue
		}
		if s[i] == ' ' {
			return 0
		}
		b = b*10 + int(s[i]) - '0'
		if int(s[i])-'0' > 9 || int(s[i])-'0' < 0 {
			return 0
		}
	}
	if d > 1 {
		return 0
	}
	if c == 1 {
		b *= -1
	}
	return b
}

func Tostr(n int) []byte {
	var str []byte
	if n == 0 {
		str = append(str, byte(0+'0'))
		return str
	}
	if n < 0 {
		n *= -1
		str = append(str, byte('-'))
	}
	temp := n
	count := 1
	for i := 0; temp > 0; i++ {
		if temp == 0 {
			break
		} else if temp <= 9 {
			break
		}
		temp /= 10  // 57 // 5 // 0
		count *= 10 // 10 // 100 // 100
	}
	temp = n % 10
	for ok := true; ok; ok = (n != 0) {
		pop := n / count
		n %= count
		count /= 10
		str = append(str, byte(pop+'0'))
	}
	if temp == 0 {
		str = append(str, byte(0+'0'))
	}
	return str
}

func IsNumeric(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || s[i] == 45 {
			count++
		}
	}
	return count == len(s)
}

func applyFunction(f func(string) int, s string) int {
	ans := f(s)
	return ans
}

func main() {
	arg := os.Args
	if len(arg) != 4 {
		return
	}
	if IsNumeric(arg[1]) && IsNumeric(arg[3]) {
		if arg[1] == "-9223372036854775809" || arg[3] == "-9223372036854775809" {
			return
		}
		a := applyFunction(Atoi, arg[1])
		b := applyFunction(Atoi, arg[3])
		ans := 0
		if arg[2] == "+" {
			ans = a + b
			if (ans > a) == (b > 0) && ans < 9223372036854775807 {
				os.Stdout.Write(Tostr(ans))
				os.Stdout.WriteString("\n")
			} else {
				return
			}
		} else if arg[2] == "-" {
			ans = a - b
			if ans > a && a < 0 && b > 0 {
				return
			}
			if (a < 0 && b > 0 && ans < a) || (ans < a && a > 0) {
				os.Stdout.Write(Tostr(ans))
				os.Stdout.WriteString("\n")
			} else {
				return
			}
		} else if arg[2] == "*" {
			ans = a * b
			if a == 0 || (ans/a == b) {
				os.Stdout.Write(Tostr(ans))
				os.Stdout.WriteString("\n")
			} else {
				return
			}
		} else if arg[2] == "/" {
			if b == 0 {
				str := []byte("No division by 0")
				os.Stdout.Write(str)
				os.Stdout.WriteString("\n")
			} else {
				ans = a / b
				os.Stdout.Write(Tostr(ans))
				os.Stdout.WriteString("\n")
			}
		} else if arg[2] == "%" {
			if b == 0 {
				str := []byte("No modulo by 0")
				os.Stdout.Write(str)
				os.Stdout.WriteString("\n")
			} else {
				ans = a % b
				os.Stdout.Write(Tostr(ans))
				os.Stdout.WriteString("\n")
			}
		}
	} else {
		return
	}
}
