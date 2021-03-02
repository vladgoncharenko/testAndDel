package recursia

func TryFact(v int) int {
	if v == 1 {
		return v
	} else {
		return v * TryFact(v-1)
	}
}
