package piscine

func StrRev(s string) string {

	origS := []byte(s)

	revS := []byte(s)

	count := 0

	for index := range origS {

		count++

		_ = index

	}

	for index := 0; index < count; index++ {

		revS[index] = origS[count-index-1]

	}

	return string(revS)

}
