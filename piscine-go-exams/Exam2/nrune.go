package piscine

func NRune(s string, n int) rune {
	ar := []rune(s) // create rune array from s length
	if n < 0 {      // failsafe to make sure - doesnt break the code
		return 0
	} else {
		for i := 0; i < n; i++ { // count up to n length
			if len(s) < n { // if string is shorter than array
				return 0 // gives us 0, failsafe
			}
			return ar[n-1] // otherwise gives us the letter, n-1 because array starts from 0
		}
		return 0 // otherwise gives us 0 again
	}
}

/*
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
*/
