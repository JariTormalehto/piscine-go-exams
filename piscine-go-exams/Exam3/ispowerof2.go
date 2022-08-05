package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func ispowerof2(n int) string {
	var i = 2
	if n == 1 || n == 2 {
		return "true"
	}

	for {
		if i == n {
			return "true"
		}

		if i > n {
			return "false"
		}

		i *= 2
	}
}

func main() {

	args := os.Args
	var otsi int

	args = os.Args

	if len(args) == 2 {
		otsi, _ = strconv.Atoi(args[1])
		ans := ispowerof2(otsi)
		for _, x := range ans {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	} else {
	}
}
