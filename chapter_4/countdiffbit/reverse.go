package countdiffbit

func Reverse(array *[6]int) {
	for i, j := 0, len(array) - 1; i < j;  {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		i++
		j--
	}
}
