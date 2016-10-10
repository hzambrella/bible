package popcount

import "testing"



func BenchPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(666)
	}
}
// !go test -bench='BenchPopCount'
/*testing: warning: no tests to run
PASS
ok		bible/2/2.6/popcount	0.006s

Press ENTER or type command to continue
*/

 //  homework 2.3
func BenchPopCountCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountCircle(666)
	}
}
// !go test -bench='BenchPopCountCircle'

/*
testing: warning: no tests to run
PASS
ok		bible/2/2.6/popcount	0.005s
*/

