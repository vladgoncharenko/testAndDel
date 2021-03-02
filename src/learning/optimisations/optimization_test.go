package optimisations

import "testing"

func BenchmarkAddToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddToSlice()
	}
}

func BenchmarkAddToSliceWithLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddToSliceWithLen()
	}
}

func BenchmarkAddToSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddToSliceWithCap()
	}
}
