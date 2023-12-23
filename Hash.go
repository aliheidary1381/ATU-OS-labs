package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

func main() {
	fmt.Print("Enter something: ")
	var userInput string
	fmt.Scanln(&userInput)
	num, err := strconv.ParseUint(userInput, 10, 32)
	if err != nil {
		return
	}
	num32 := uint32(num)

	// test := uint32(4)
	byteSlice := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteSlice, num32)
	hasher := fnv.New32()
	hasher.Write(byteSlice)
	hashTest := hasher.Sum32()

	maxUint32 := math.MaxUint32
	rangeForEachPart := uint32(maxUint32 / 4)

	part1 := uint32(0)
	part2 := part1 + rangeForEachPart
	part3 := part2 + rangeForEachPart
	part4 := part3 + rangeForEachPart

	x1 := CalculateHash(part1, part2, hashTest)
	x2 := CalculateHash(part2, part3, hashTest)
	x3 := CalculateHash(part3, part4, hashTest)
	x4 := CalculateHash(part4, uint32(maxUint32), hashTest)

	_ = x1
	_ = x2
	_ = x3
	_ = x4

	// fmt.Println(x1)
	// fmt.Println(x2)
	// fmt.Println(x3)
	// fmt.Println(x4)
}

func CalculateHash(start, end, hash uint32) chan int {

	ch := make(chan int)
	go func() {
		for x := start; x < end; x++ {
			temp := x
			byteSlice := make([]byte, 4)
			binary.LittleEndian.PutUint32(byteSlice, x)
			hasher := fnv.New32()
			hasher.Write(byteSlice)
			hashValue := hasher.Sum32()
			if hashValue == hash {
				fmt.Printf(" %d \n", temp)
				ch <- int(x)
				// return x
			}
		}
	}()

	return ch
}
