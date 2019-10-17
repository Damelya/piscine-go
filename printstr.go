package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	nstring := []rune(str)
	for i := range nstring {
		z01.PrintRune(nstring[i])
	}
}
