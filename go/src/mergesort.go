package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func mergesort(data []int, low int, high int) {
	var middle = (low + high) / 2
	if (high - low) > 0 {
		mergesort(data, low, middle)
		mergesort(data, middle+1, high)
		merge(data, low, middle, high)
	}
}

func merge(data []int, low int, middle int, high int) {
	// copy the lower half and higher half into separate buffers.
	buf1 := []int{}
	buf2 := []int{}

	//copy the first half
	for i := low; i <= middle; i++ {
		buf1 = append(buf1, data[i])
	}

	// copy the second half.
	for i := middle + 1; i <= high; i++ {
		buf2 = append(buf2, data[i])
	}

	//compare until one of the buffer runs out.
	i := low
	for !(len(buf1) == 0 || len(buf2) == 0) {
		if buf1[0] < buf2[0] {
			data[i] = buf1[0]
			buf1 = buf1[1:]
		} else {
			data[i] = buf2[0]
			buf2 = buf2[1:]
		}
		i++
	}

	for len(buf1) > 0 {
		data[i] = buf1[0]
		i++
		buf1 = buf1[1:] // pop the top element
	}

	for len(buf2) > 0 {
		data[i] = buf2[0]
		i++
		buf2 = buf2[1:] // pop the top element
	}

}

func main() {
	var (
		data      = []int{}
		inputSize = 10
	)
	for i := 0; i < inputSize; i++ {
		data = append(data, rand.Intn(100))
	}
	fmt.Println(data)
	mergesort(data, 0, len(data)-1)
	if sort.IntsAreSorted(data) {
		fmt.Println(data)
	} else {
		fmt.Println("data is not sorted")
	}

}

/*
 * Learnings
 * sort package has lot more functionalities. have to learn more about that.
 * generating random ints using math/rand package
 * deleting a slice while iterating over it is tricky.
 * for popping, I used l = l[1:]. for deleting an item at index i, I can use
 * l = append(l[:i], l[i+1:]...) ... is important here. It didnt work without that.
 * No while loop in Go.
 * No assert in Go.
 * Sometimes I forget to increment the counters in loop and also add
 * stopping condition for recursive loops.
 */
