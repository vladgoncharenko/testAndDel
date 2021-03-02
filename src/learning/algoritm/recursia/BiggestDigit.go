package recursia

func BiggestDigit(data []int) int {
	if len(data) == 1 {
		return data[0]
	}
	if data[0] > data[1] {
		data[1] = data[0]
	}
	return BiggestDigit(data[1:])
}
