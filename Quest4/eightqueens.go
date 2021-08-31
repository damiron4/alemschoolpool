package piscine

import "github.com/01-edu/z01"

var (
	cboard [8][8]bool
	attack [8][8]int
)

func PlaceQueen(i, j int) {
	cboard[i][j] = true
	for c := 0; c < 8; c++ {
		attack[i][c]++
	}
	for r := 0; r < 8; r++ {
		attack[r][j]++
	}
	attack[i][j] -= 2
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				attack[r][c]++
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				attack[r][c]++
			}
		}
	}
}

func RemoveQueen(i, j int) {
	cboard[i][j] = false
	for c := 0; c < 8; c++ {
		attack[i][c]--
	}
	for r := 0; r < 8; r++ {
		attack[r][j]--
	}
	attack[i][j] += 2
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				attack[r][c]--
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				attack[r][c]--
			}
		}
	}
}

func Solve(r int) {
	if r == 7 {
		for c, cnt := range attack[7] {
			if cnt == 0 {
				PlaceQueen(r, c)
				Table(cboard)
				RemoveQueen(r, c)
			}
		}
	} else {
		for c, cnt := range attack[r] {
			if cnt == 0 {
				PlaceQueen(r, c)
				Solve(r + 1)
				RemoveQueen(r, c)
			}
		}
	}
}

func Table(data [8][8]bool) {
	str := "12345678"
	for i := 0; i < 8; i++ {
		row := make([]int, 8)
		for j := 0; j < 8; j++ {
			if data[i][j] {
				row[j] = j
				z01.PrintRune(rune(str[row[j]]))
			}
		}
	}
	z01.PrintRune('\n')
}

func EightQueens() {
	Solve(0)
}
