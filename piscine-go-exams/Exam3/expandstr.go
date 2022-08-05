package main

import "os"

func main() {
	Expandstr()
}

func Expandstr() {
	if len(os.Args) != 2 {
		return
	}
	answer := ""
	isWord := false
	flag := 0
	for _, letter := range os.Args[1] {
		if isValid(letter) {
			isWord = true
			flag = 1
		} else if letter == ' ' {
			isWord = false
			if flag == 1 {
				answer += "   "
				flag = 0
			}
		}
		if isWord {
			answer += string(letter)
		}
	}
	if answer[len(answer)-1] == ' ' {
		answer = answer[:len(answer)-3]
	}
	os.Stdout.WriteString(answer + "\n")
}

func isValid(v rune) bool {
	if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
		return true
	}
	return false
}
