package piscine

func Swap(a *int, b *int) {
	c := *a
	*a = *a
	*b = c
}
