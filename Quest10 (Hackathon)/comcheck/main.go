package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, word := range arg {
		if word == "galaxy" || word == "01" || word == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
