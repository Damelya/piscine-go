package piscine

func StrLen(str string) int {
	len := 0
	nstring := []rune(str)
	for i := range nstring {
		len++
		i = i
	}
	return len
}
