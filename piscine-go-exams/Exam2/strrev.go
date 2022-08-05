package main

import "fmt"

/*
the idea behind -- reverse = string(u) + reverse -- is:
 1) reverse = first letter from range s(main function) + reverse (H)
 2) reverse = second letter from range s(main function)+ reverse (eH)
 3) reverse = third letter from range s(main function) + reverse (leH)

 main function first means it adds the next letter and + reverse after that, "reverse" stores previous
 values
*/
func StrRev(s string) string {
	var reverse string    // create variable reverse type string without a value
	for _, u := range s { // takes the length of the string s
		reverse = string(u) + reverse // give variable reverse a value of range + reverse
	}
	return reverse // print out the reverse value
}

/*
var reverse string
for i := 0; i <= len(s)-1; i++ {
	reverse = string(s[i]) + reverse
}
return reverse
*/

/*
	func StrRev(s string) string {

	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
} */

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
