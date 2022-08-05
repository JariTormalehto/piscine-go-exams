package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	answerSlice := []string{}
	var appendString string
	if len(os.Args) == 1 || len(os.Args) > 2 {
		z01.PrintRune('\n')
	} else {
		inputSlice := []rune(os.Args[1])
		for i := 0; i <= len(inputSlice); i++ {
			if i == len(inputSlice) {
				answerSlice = append(answerSlice, appendString)
				break
			}
			if inputSlice[i] != 32 && inputSlice[i] != 9 && inputSlice[i] != 10 {
				appendString += string(inputSlice[i])
			}
			if (len(appendString) != 0) && (inputSlice[i] == 32 || inputSlice[i] == 9 || inputSlice[i] == 10) {
				answerSlice = append(answerSlice, appendString)
				appendString = ""
			}
		}
		for a := 0; a < len(answerSlice); a++ {
			printString := answerSlice[a]
			printSlice := []rune(printString)
			for j := 0; j < len(printSlice); j++ {
				z01.PrintRune(printSlice[j])
			}
			if a < len(answerSlice)-1 && answerSlice[a] != "harder" {
				z01.PrintRune(32)
			}
		}
		z01.PrintRune('\n')
	}
}
