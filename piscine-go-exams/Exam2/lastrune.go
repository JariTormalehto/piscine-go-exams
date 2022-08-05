package piscine

func LastRune(s string) rune {
	ar := []rune(s)
	return ar[len(s)-1]
}

/*
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
*/
