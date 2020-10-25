package learning

func search(sortedArray []int, element int) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		middle := ((end - init) >> 1) + init
		if sortedArray[middle] == element {
			return middle
		}
	}
	return 0
}
