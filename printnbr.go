package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	if n < 0 {

		z01.PrintRune('-')

	}

	PrintOne(n)

}

func PrintOne(numb int) {

	a := '0'

	if numb == 0 {

		z01.PrintRune(a)

		return

	}

	for b := 1; b <= numb%10; b++ {

		a++

	}

	for b := -1; b >= numb%10; b-- {

		a++

	}

	if numb/10 != 0 {

		PrintOne(numb / 10)

	}

	z01.PrintRune(a)

	return

}
