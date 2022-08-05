package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	for i, v := range base {
		for j, w := range base {
			if w == v && i != j || v == '-' || v == '+' {
				return 0
			}
		}
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += Index_(base, string(s[i])) * Pow(len(base), len(s)-i-1)
	}
	return sum
}

func Index_(base, nbr string) int {
	for i := range base {
		if string(base[i]) == nbr {
			return i
		}
	}
	return 0
}

func Pow(nbr, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * nbr
	}
	return res
}
