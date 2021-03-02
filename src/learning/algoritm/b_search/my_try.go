package b_search

func SearchB(target int, data []int) int {
	start, end := 0, len(data)-1

	for start <= end {
		middle := (start + end) / 2
		if data[middle] == target {
			return middle
		}

		if target > data[middle] {
			start = middle
		} else {
			end = middle
		}
	}
	return -1

}
