package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <= 999; a++ {
		b = a / 100
		c = a / 10
		d = a % 10
		if b < c < d {
			z01.PrintRune(a)
		}
	}
}
