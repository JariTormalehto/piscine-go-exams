package main

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, s := range str { // takes the length of string provided in main function
		z01.PrintRune(s) // prints it out one by one
	}
}

func main() {
	PrintStr("Hello World!")
}
