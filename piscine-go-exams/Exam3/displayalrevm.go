package main

import (
	"github.com/01-edu/z01"
)

func main() {

	list := "aBcDeFgHiJkLmNoPqRsTuVxXyZ"

	for i := range list {
		z01.PrintRune(rune(list[i]))
	}
	z01.PrintRune('\n')

}
