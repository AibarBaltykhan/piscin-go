package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 1; a <= 99; a++ {
		c := a / 10
		d := a % 10
		if c < d {
			z01.PrintRune(rune('0'))
			z01.PrintRune(rune(a + 60))
		}
	}

	for a := 100; a <= 999; a++ {
		b := a / 100
		c := a / 10
		d := a % 10
		if (b < c) && (c < d) {

			z01.PrintRune(rune(a + 60))
		}
	}
}
