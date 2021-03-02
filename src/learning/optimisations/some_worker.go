package optimisations

func AddToSlice() {
	var data []string

	for i := 0; i < 100; i++ {
		data = append(data, "some_text")
	}
}

func AddToSliceWithLen() {
	var data = make([]string, 20)

	for i := 0; i < 100; i++ {
		data = append(data, "some_text")
	}

}

func AddToSliceWithCap() {
	var data = make([]string, 0, 20)

	for i := 0; i < 100; i++ {
		data = append(data, "some_text")
	}
}
