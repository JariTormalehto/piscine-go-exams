package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		arguments := os.Args[1:]
		count := 0

		for i := range arguments {
			count++
			i += i
		}
		z01.PrintRune(rune('0' + count))
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
