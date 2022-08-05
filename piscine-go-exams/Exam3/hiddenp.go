package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	str1 := os.Args[1]
	str2 := os.Args[2]

	var match string
	count := 0

	if len(os.Args) == 3 {
		for i := 0; i < len(str1); i++ {
			for j := count; j < len(str2); j++ {
				if str1[i] == str2[j] {
					match += string(str1[i])
					count = j + 1
					break
				}
			}
		}
		if match == str1 {
			z01.PrintRune('1')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		}
	}

}
