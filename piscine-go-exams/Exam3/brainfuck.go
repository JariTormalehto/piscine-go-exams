package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		Brainfuck(ar[1])
	} else {
		printer("\n")
	}
}

func Brainfuck(s string) {
	b := [2048]byte{}
	x := 0
	p := &b[x]
	depth := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '>':
			x++
		case '<':
			x--
		case '-':
			*p--
		case '+':
			*p++
		case '.':
			printer(string(*p))
		case '[':
			depth++
			if *p == 0 {
				oldDepth := depth - 1
				for oldDepth != depth {
					i++
					if s[i] == '[' {
						depth++
					} else if s[i] == ']' {
						depth--
					}
				}
			}
		case ']':
			depth--
			if *p != 0 {
				oldDepth := depth + 1
				for oldDepth != depth {
					i--
					if s[i] == '[' {
						depth++
					} else if s[i] == ']' {
						depth--
					}
				}
			}
		}
		p = &b[x]
	}
}

func printer(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
