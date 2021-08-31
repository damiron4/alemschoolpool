package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Opfile() {
	content, _ := ioutil.ReadFile("quest8.txt")
	fmt.Print(string(content))
}

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("File name missing")
	} else if len(arg) > 2 {
		fmt.Println("Too many arguments")
	} else {
		Opfile()
	}
}
