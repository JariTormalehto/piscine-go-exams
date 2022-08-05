package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	bruh1, _ := strconv.Atoi(arg[0])
	bruh2, _ := strconv.Atoi(arg[1])

	result := gcd(bruh1, bruh2)
	fmt.Print(result)
	fmt.Println()

}
func gcd(bruh1, bruh2 int) int {
	for bruh2 != 0 {
		bruh1, bruh2 = bruh2, bruh1%bruh2
	}
	return bruh1
}
