package piscine

func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	} else if n < 25 {
		return n * RecursiveFactorial(n-1)
	}
	return 0
}
