package q1

import (
	"os"
	"testing"
)

var x, y []uint64

func TestMain(m *testing.M) {
	var A [l]uint64
	x = A[:]
	y = A[:]
	code := m.Run()
	os.Exit(code)
}

func BenchmarkMethodA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Axpy1(2, x, y)
	}
}

func BenchmarkMethodB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Axpy2(2, x[0:], y[0:])
		Axpy2(2, x[1:], y[1:])
	}
}

func BenchmarkMethodC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Axpy3(2, x, y)
	}
}

func BenchmarkMethodD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Axpy4(2, x, y)
	}
}
