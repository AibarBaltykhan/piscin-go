package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	int a, b, c
	b = a / 100
	c = a / 10
	d = a % 10
	for a := 0; a <= 999; a++ {
		if b < c < d {
			z01.PrintRune(a)
		}
	}
}
