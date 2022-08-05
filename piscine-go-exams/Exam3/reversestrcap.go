package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Capitalize(s string) string {
	a := 0
	array := []rune(s)
	for range s {
		a++
	}
	for i := 0; i < a; i++ {
		if array[i] >= 'A' && array[i] <= 'Z' {
			array[i] = array[i] + 32
		}
		if array[0] >= 'a' && array[0] <= 'z' {
			array[0] = array[0] - 32
		}
		if i > 0 {
			if array[i-1] > 'Z' && array[i-1] < 'a' || array[i-1] < 'A' && array[i-1] > '9' || array[i-1] < '0' && array[i-1] != '\'' || array[i-1] > 'z' {
				if array[i] >= 'a' && array[i] <= 'z' {
					array[i] = array[i] - 32
				}
			}
		}
	}
	return string(array)
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		return
	}
	for _, str := range arguments {
		var revresedString []byte
		for i := range str {
			revresedString = append(revresedString, str[len(str)-1-i])
		}
		capString := Capitalize(string(revresedString))
		for i := len(capString) - 1; i >= 0; i-- {
			z01.PrintRune(rune(capString[i]))
		}
		z01.PrintRune('\n')
	}
}
