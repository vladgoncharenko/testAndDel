package recursia

import "fmt"

func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	left, right := 0, len(data)-1

	pivot := len(data) % 2

	data[pivot], data[right] = data[right], data[pivot]

	fmt.Println("pivot ", data[pivot])
	fmt.Println("right ", data[right])
	fmt.Println("left ", data[left])

	for i := range data {
		fmt.Println("i ", data[i])
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	fmt.Println("_____________")
	fmt.Println(data)
	fmt.Println("right ", data[right])
	fmt.Println("left ", data[left])
	fmt.Println("_____________")

	QuickSort(data[:left])

	QuickSort(data[left+1:])

	return data
}
