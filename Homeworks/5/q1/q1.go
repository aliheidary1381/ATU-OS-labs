package q1

import "sync"

const l = 128 * 1024 * 1024 * 8 / 64

func Axpy1(a uint64, x []uint64, y []uint64) {
	for i := 0; i < l; i++ {
		y[i] += a * x[i]
	}
}

func Axpy2(a uint64, x []uint64, y []uint64) {
	for i := 0; i < l; i += 2 {
		y[i] += a * x[i]
	}
}

func Axpy3(a uint64, x []uint64, y []uint64) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		Axpy1(a, x[:l/2], y[:l/2])
		wg.Done()
	}()
	go func() {
		Axpy1(a, x[l/2:], y[l/2:])
		wg.Done()
	}()
	wg.Wait()
}

func Axpy4(a uint64, x []uint64, y []uint64) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		Axpy2(a, x[0:], y[0:])
		wg.Done()
	}()
	go func() {
		Axpy2(a, x[1:], y[1:])
		wg.Done()
	}()
	wg.Wait()
}
