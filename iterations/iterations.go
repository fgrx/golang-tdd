package iterations

const loopNumber = 5

func Repeat(letter string) string {
	var res string

	for i := 0; i < loopNumber; i++ {
		res += letter
	}

	return res
}