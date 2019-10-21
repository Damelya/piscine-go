package piscine

func RecursivePower(nb int, power int) int {

	if power > 0 {

		if power == 1 {

			return nb

		} else {

			return nb * RecursivePower(nb, power-1)

		}

	} else if power == 0 {

		return 1

	}

	return 0

}
