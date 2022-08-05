package piscine

func Rot14(s string) string {
	result := ""
	for _, res := range s {
		result += string(rot14(res))
	}
	return result
}

func rot14(b rune) rune {
	if b >= 'A' && b < 'M' || b >= 'a' && b < 'm' {
		return b + 14
	}
	if b >= 'M' && b <= 'Z' || b >= 'm' && b <= 'z' {
		return b - 12
	}
	return b
}
