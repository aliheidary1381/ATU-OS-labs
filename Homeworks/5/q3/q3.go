package q3

import "sync"

const C = 1_073_741_824

func PP1(a []uint) {
	for i := 0; i < C; i++ {
		a[0]++
	}
	for i := 0; i < C; i++ {
		a[1]++
	}
}

func PP2(a []uint) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < C; i++ {
			a[0]++
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < C; i++ {
			a[1]++
		}
		wg.Done()
	}()
	wg.Wait()
}
