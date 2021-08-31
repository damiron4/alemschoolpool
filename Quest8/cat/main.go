package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			file, err := ioutil.ReadFile(arg)
			if err != nil {
				z01.PrintRune('E')
				z01.PrintRune('R')
				z01.PrintRune('R')
				z01.PrintRune('O')
				z01.PrintRune('R')
				z01.PrintRune(':')
				z01.PrintRune(' ')
				io.WriteString(os.Stdout, err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			os.Stdout.Write(file)
		}
	}
}
