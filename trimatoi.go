package piscine

func TrimAtoi(s string) int {

	i := 0

	for _, letter := range s {

		if letter >= '0' && letter <= '9' {

			if letter == '0' {

				i += 0

				i *= 10

			} else if letter == '1' {

				i += 1

				i *= 10

			} else if letter == '2' {

				i += 2

				i *= 10

			} else if letter == '3' {

				i += 3

				i *= 10

			} else if letter == '4' {

				i += 4

				i *= 10

			} else if letter == '5' {

				i += 5

				i *= 10

			} else if letter == '6' {

				i += 6

				i *= 10

			} else if letter == '7' {

				i += 7

				i *= 10

			} else if letter == '8' {

				i += 8

				i *= 10

			} else if letter == '9' {

				i += 9

				i *= 10

			}

		}

	}

	number := i / 10

	minindex := 100

	lindex := 0

	for index, letter := range s {

		if letter == '-' {

			minindex = index

		}

		if letter == '1' || letter == '2' || letter == '3' || letter == '4' || letter == '5' || letter == '6' || letter == '7' || letter == '8' || letter == '9' {

			lindex = index

			if lindex > minindex {

				return -number

			} else if lindex > 0 {

				return number

			}

		}

	}

	return number

}
