package main

import (
	"fmt"
	"runtime"
	"sync"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSortParallel(arr []int, depth int, wg *sync.WaitGroup) []int {
	defer wg.Done()

	if len(arr) <= 1 {
		return arr
	}

	if depth <= 0 {
		return mergeSort(arr)
	}

	mid := len(arr) / 2

	var left, right []int
	var wgInner sync.WaitGroup

	wgInner.Add(2)

	go func() {
		left = mergeSortParallel(arr[:mid], depth-1, &wgInner)
	}()

	go func() {
		right = mergeSortParallel(arr[mid:], depth-1, &wgInner)
	}()

	wgInner.Wait()

	return merge(left, right)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted:", arr)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		arr = mergeSortParallel(arr, 2, &wg)
	}()

	wg.Wait()

	fmt.Println("Sorted:", arr)
}
