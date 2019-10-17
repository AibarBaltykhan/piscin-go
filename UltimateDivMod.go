package piscine

func UltimateDivMod(a *int, b *int) {
	*x := 0
	*y := 0
	*x = *a
	*y = *b
	*a = *x / *y
	*b = *x % *y
}
