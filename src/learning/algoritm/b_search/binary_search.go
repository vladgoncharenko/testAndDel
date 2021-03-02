package b_search

import "fmt"

func Search(sortedArray []int, element int) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		fmt.Println(init, "  ", end)
		middle := ((end - init) / 2) + init
		fmt.Println(middle)
		if sortedArray[middle] == element {
			return middle
		}
		if sortedArray[middle] < element {
			init = middle + 1
		} else {
			end = middle - 1
			println()
		}
	}
	return -1
}
