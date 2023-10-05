package q2

import (
	"os"
	"testing"
)

var (
	B [][l]uint
	A [][l]uint
)

func TestMain(m *testing.M) {
	var a [l][l]uint
	A = a[:]
	B = a[:]
	code := m.Run()
	os.Exit(code)
}

func BenchmarkMethodA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GeMM1(A, B)
	}
}

func BenchmarkMethodB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GeMM2(A, B)
	}
}
