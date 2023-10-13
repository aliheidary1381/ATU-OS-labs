package q3

import (
	"testing"
)

var A [2]uint

func BenchmarkMethodA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PP1(A[:])
	}
}

func BenchmarkMethodB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PP2(A[:])
	}
}
