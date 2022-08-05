package main

import (
	"os"

	"github.com/01-edu/z01"

	"strconv"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		intn, err := strconv.Atoi(arg[0])
		if err == nil {
			if intn < 0 {
				z01.PrintRune(10)
				return
			}
			for i := 1; i < 10; i++ {
				bruh(i)
				z01.PrintRune(' ')
				z01.PrintRune('x')
				z01.PrintRune(' ')
				bruh(intn)
				z01.PrintRune(' ')
				z01.PrintRune('=')
				z01.PrintRune(' ')
				bruh(i * intn)
				z01.PrintRune(10)
			}
		}
	}
}

func bruh(lol int) {
	if lol >= 10 {
		bruh(lol / 10)
	}
	z01.PrintRune(rune(lol%10 + 48))
}
