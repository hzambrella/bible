package popcount

import "testing"

// !go test -bench='BenchPopCount'
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(666)
	}
}

func BenchmarkPopCountCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountCircle(666)
	}

}
