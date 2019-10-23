package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	var nb []int

	count := 0

	for n > 0 {

		nb = append(nb, n%10)

		count++

		n /= 10

	}

	if count == 0 && n == 0 {

		z01.PrintRune('0')

	}

	for j := 0; j < count-1; j++ {

		for k := 0; k < count-j-1; k++ {

			if nb[k] > nb[k+1] {

				nb[k], nb[k+1] = nb[k+1], nb[k]

			}

		}

	}

	for l := 0; l < count; l++ {

		z01.PrintRune(rune(nb[l]) + '0')

	}

}
