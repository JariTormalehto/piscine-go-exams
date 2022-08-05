package main

import (
	"os"
)

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]
		var firstNum int
		var secondNum int
		var operator string
		IsfirstNum := false
		IsSecondNum := false
		if string(args[0][0]) == "-" {
			firstNum, IsfirstNum = findNumber(args[0][1:], args[0])
		} else if string(args[0][0]) != "-" {
			firstNum, IsfirstNum = findNumber(args[0], args[0])
		}
		if string(args[2][0]) != "-" {
			secondNum, IsSecondNum = findNumber(string(args[2]), string(args[2]))
		} else if string(args[2][0]) == "-" {
			secondNum, IsSecondNum = findNumber(args[2][1:], args[2])
		}
		if IsOperator(args[1]) && overFlow(firstNum) && overFlow(secondNum) {
			operator = args[1]
			if IsSecondNum && IsfirstNum {
				switch operator {
				case "+":
					os.Stdout.Write(intToByte(firstNum + secondNum))
				case "-":
					os.Stdout.Write(intToByte(firstNum - secondNum))
				case "*":
					os.Stdout.Write(intToByte(firstNum * secondNum))
				case "/":
					if secondNum == 0 {
						strToByte("No division by 0\n")
					} else {
						os.Stdout.Write(intToByte(firstNum / secondNum))
					}
				case "%":
					if secondNum == 0 {
						strToByte("No modulo by 0\n")
					} else {
						os.Stdout.Write(intToByte(firstNum % secondNum))
					}
				}
			}
		}
	}
}

func intToByte(num int) []byte {
	var lists []byte
	negNum := false
	if num == 0 {
		lists = append(lists, 48)
		lists = append(lists, "\n"[0])
		return lists
	}
	if num < 0 {
		negNum = true
		num = num * -1
	}
	for 0 < num {
		lists = append([]byte{byte(num%10 + 48)}, lists...)
		num = num / 10
	}
	if negNum {
		lists = append([]byte{byte(45)}, lists...)
	}
	lists = append(lists, "\n"[0])
	return lists
}

func strToByte(s string) {
	var lists []byte
	for i := range s {
		lists = append(lists, s[i])
	}
	os.Stdout.Write(lists)
}

func findNumber(s string, s2 string) (int, bool) {
	if IsNumeric(s) {
		return Atoi(s2), true
	}
	return 0, false
}

func overFlow(num int) bool {
	if num < 9223372036854775807 && -9223372036854775807 < num {
		return true
	}
	return false
}

func IsOperator(s string) bool {
	if s == "+" || s == "-" || s == "/" || s == "%" || s == "*" {
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	for _, symbol := range s {
		if symbol >= 48 && symbol <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	newInt := 0
	var number int
	plussorminus := 1
	for posnum, symbol := range s {
		if posnum == 0 && symbol == 43 {
			continue
		} else if posnum == 0 && symbol == 45 {
			plussorminus = -1
		} else if symbol >= 48 && symbol <= 57 {
			for i := '0'; symbol > i; i++ {
				number++
			}
			newInt = (newInt * 10) + number
			number = 0
		} else {
			return 0
		}
	}
	return newInt * plussorminus
}
