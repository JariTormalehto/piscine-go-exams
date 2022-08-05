package piscine

func StrLen(s string) int {
	a := 0        // variable a type integer with starting value of a
	for range s { // takes the string s length
		a++ // for every letter(length) keeps adding +1 to a value
	}
	return (a) // end value of string s in integer length
}

/*
func StrLen(s string) int {
	return len([]rune(s))
}

*/
