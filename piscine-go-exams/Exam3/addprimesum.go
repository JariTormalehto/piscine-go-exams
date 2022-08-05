package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsPrime(nb int) bool {
	result := true

	if nb <= 1 {
		result = false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			result = false
			break
		}
	}
	return result
}

func Atoi(s string) int {
	symbol := 0
	isNeg := false
	slice := []rune(s)

	if len(s) == 0 {
		return 0
	}

	if slice[0] == '-' {
		slice = slice[1:]
		isNeg = true
	} else if slice[0] == '+' {
		slice = slice[1:]
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return 0
		}
	}
	for _, inp := range slice {
		symbol = (symbol * 10) + int(inp-'0')
	}
	if isNeg {
		symbol = symbol * (-1)
	}
	return symbol
}
func PrintNbr(n int) {
	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		argument := Atoi(os.Args[1])

		if argument < 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		} else {
			result := 0
			for ; argument >= 0; argument-- {
				if IsPrime(argument) {
					result += argument
				}
			}
			PrintNbr(result)
			z01.PrintRune('\n')
		}
	}
}
