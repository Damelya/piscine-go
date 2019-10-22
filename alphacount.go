package piscine

func AlphaCount(str string) int {

	count := 0

	for _, letter := range str {

		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {

			count++

		}

	}

	return count

}
