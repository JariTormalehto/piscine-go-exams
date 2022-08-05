package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	myMaximum := a[0]
	for _, value := range a {
		if value > myMaximum {
			myMaximum = value
		}
	}
	return myMaximum
}
