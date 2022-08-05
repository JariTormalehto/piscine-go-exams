package piscine

func Swap(a *int, b *int) {
	tempA := *a // create a variable tempA which has the value of pointer a
	tempB := *b // create a variable tempB which has the value of pointer b
	*a = tempB  // give pointer a the value of tempb, thus reversing a into b
	*b = tempA  // give pointer b the value of tempa, thus reversing b into a
}
