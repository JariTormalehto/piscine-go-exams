package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		for _, v := range os.Args[1] {
			if v == []rune(os.Args[2])[0] {
				v = []rune(os.Args[3])[0]
			}
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
