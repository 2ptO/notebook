package minheap

import "errors"

//MinHeap ...exported for testing
type MinHeap struct {
	size, capacity int
	data           []int
}

//ErrEmptyHeap ...Returned when heap is Empty
var ErrEmptyHeap = errors.New("Empty Heap")

func (mHeap MinHeap) peek() (int, error) {
	if mHeap.size == 0 {
		return 0, ErrEmptyHeap
	}
	return mHeap.data[0], nil
}

func (mHeap *MinHeap) add(newVal int) {
	// mHeap.ensureEnoughCapacity()
	mHeap.data[mHeap.size] = newVal
	mHeap.bubbleUp(mHeap.size)
	mHeap.size++
}

func (mHeap *MinHeap) bubbleUp(index int) {
	// Remember: Recursion with array is dangerous.
	// change this into while loop instead.
	for mHeap.hasParent(index) &&
		(mHeap.data[index] < mHeap.data[mHeap.getParentIndex(index)]) {
		mHeap.swap(index, mHeap.getParentIndex(index))
		index = mHeap.getParentIndex(index)
	}
}

func (mHeap MinHeap) hasParent(index int) bool {
	if index == 0 {
		return false
	}
	return true
}

func (mHeap MinHeap) getParentIndex(childIndex int) int {
	// not feeling good about this. It gives a chance of error when index is 0
	return (childIndex - 1) / 2
}

func (mHeap *MinHeap) swap(from, to int) {
	tmp := mHeap.data[from]
	mHeap.data[from] = mHeap.data[to]
	mHeap.data[to] = tmp
}

func (mHeap *MinHeap) pop() (int, error) {
	root, err := mHeap.peek()
	if err != nil {
		return 0, err
	}
	mHeap.swap(0, (mHeap.size - 1))
	mHeap.size--
	mHeap.bubbleDown(0)
	return root, nil
}

func (mHeap *MinHeap) bubbleDown(parentIndex int) {
	for mHeap.hasLeftChild(parentIndex) {
		smallerChildIndex := mHeap.getLeftChildIndex(parentIndex)
		if mHeap.hasRightChild(parentIndex) &&
			(mHeap.data[mHeap.getRightChildIndex(parentIndex)] < mHeap.data[smallerChildIndex]) {
			smallerChildIndex = mHeap.getRightChildIndex(parentIndex)
		}
		if mHeap.data[smallerChildIndex] < mHeap.data[parentIndex] {
			mHeap.swap(parentIndex, smallerChildIndex)
			parentIndex = smallerChildIndex
		} else {
			break
		}
	}
}

func (mHeap MinHeap) hasLeftChild(parentIndex int) bool {
	return mHeap.getLeftChildIndex(parentIndex) < mHeap.size
}

func (mHeap MinHeap) hasRightChild(parentIndex int) bool {
	return mHeap.getRightChildIndex(parentIndex) < mHeap.size
}

func (mHeap MinHeap) getRightChildIndex(parentIndex int) int {
	return (parentIndex * 2) + 2
}

func (mHeap MinHeap) getLeftChildIndex(parentIndex int) int {
	return (parentIndex * 2) + 1
}
