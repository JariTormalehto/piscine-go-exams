package piscine

func FirstRune(s string) rune {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	return first
}

/*
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
*/
