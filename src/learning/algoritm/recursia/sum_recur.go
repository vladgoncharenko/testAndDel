package recursia

func SumRecur(data []int) int {
	if len(data) == 0 {
		return 0
	}
	return data[0] + SumRecur(data[1:])
}
