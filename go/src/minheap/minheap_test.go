package minheap

import (
	"fmt"
	"sort"
	"testing"
)

func TestHeapReturnsErrorWhenEmpty(t *testing.T) {
	mHeap := MinHeap{0, 0, make([]int, 0, 10)}
	_, err := mHeap.peek()
	if err != ErrEmptyHeap {
		t.Fatal("No error returned when heap is empty")
	}
	t.Log(err)
}

func TestHeapRootIsMinimum(t *testing.T) {
	mHeap := MinHeap{0, 0, make([]int, 10)}
	testInputs := []int{10, 8, 20, 15, 4, 40, 13}
	for _, testVal := range testInputs {
		mHeap.add(testVal)
	}
	sort.Ints(testInputs)
	if val, _ := mHeap.peek(); val != testInputs[0] {
		fmt.Println(mHeap.data)
		t.Fatal("Root of the heap is not the minimum value")
	}
}

func TestHeapPopWhenEmpty(t *testing.T) {
	mHeap := MinHeap{0, 0, make([]int, 10)}
	_, err := mHeap.pop()
	if err != ErrEmptyHeap {
		t.Fatal("No error received when heap is empty")
	}
}

func TestHeapPop(t *testing.T) {
	mHeap := MinHeap{0, 0, make([]int, 10)}
	testInputs := []int{10, 8, 20, 15, 4, 40, 13}
	//add inputs to Heap
	for _, testVal := range testInputs {
		mHeap.add(testVal)
	}

	//sort testInputs so that we can pop and compare
	sort.Ints(testInputs)
	for _, expected := range testInputs {
		if actual, err := mHeap.pop(); err != nil || actual != expected {
			t.Fatalf("Actual %d, expected %d\n", actual, expected)
		}
	}
}
