package main

import (
	"fmt"
	"io"
	)

func HeapSort(b []int) {
	// Consider b to be a commplete binary tree.
	// Build heap with greatest element on the top.
	for i := (len(b)-1) / 2; i > 0; i-- {
		HeapAdjust(b, i)
	}
	// Continually Pop heap top elements and adjust remaining elements to be a heap.
	for i := len(b) - 1; i > 1; i-- {
		//swap the top element with the last element on the heap.
		b[i], b[1] = b[1], b[i]
		// Generate new heap.
		HeapAdjust(b[:i], 1)
	}
}

func HeapAdjust(b []int, i int) {
	// Siftdown i
	siftdata := b[i] //Save the element to be siftdown.
	for child := 2 * i; child <= len(b)-1; child = 2 * child {
		//Siftdown along with the bigger child element branch.
		if child < len(b)-1 && b[child] < b[child+1] {
			child++
		}
		if siftdata >= b[child] {
			break
		} else {
			b[i] = b[child] // Adjust b[child] to parent node.
			i = child       // Update i for continuing siftdown.
		}
	}
	b[i] = siftdata // Put the siftdata at the final position.
}

func IsHeap(b []int) bool { // If s is a big endian heap.
	length := len(b) - 1
	if length % 2 == 0 { // Len(s) is even number. S contains a single-branch node.
		if b[length/2] < b[length] {
			return false
		}
		for i := length/2 - 1; i >= 1; i-- {
			if b[i] < b[2*i] || b[i] < b[2*i+1] {
				return false
			}
		}
	} else { // Len(s) is odd number,none single-branch node.
		for i := length / 2; i >= 1; i-- {
			if b[i] < b[2*i] || b[i] < b[2*i+1] {
				return false
			}
		}
	}
	return true
}

func main() {
	var a int
	l := make([]int, 1)
	l[0] = 0
	for {
		_, err := fmt.Scan(&a)
		if err == io.EOF {
			break
		}
		l = append(l, a)
	}
	if !IsHeap(l) {
		HeapSort(l)
		fmt.Println(l[1:])
	} else {
		fmt.Println(l[1:], "is already a heap.")
	}
}
