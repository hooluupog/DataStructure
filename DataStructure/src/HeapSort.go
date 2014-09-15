package main

import "fmt"

func HeapSort(s string) string {
	b := make([]byte, len(s)+1)
	copy(b[1:], []byte(s))
	// Consider s to be a commplete binary tree.
	// Build heap with greatest element on the top.
	for i := len(s) / 2; i > 0; i-- {
		HeapAdjust(b, i)
	}
	// Continually Pop heap top elements and adjust remaining elements to be a heap.
	for i := len(s); i > 1; i-- {
		//swap the top element with the last element on the heap.
		b[i], b[1] = b[1], b[i]
		// Generate new heap.
		HeapAdjust(b[:i], 1)
	}
	return string(b[1:])
}

func HeapAdjust(b []byte, i int) {
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

func IsHeap(s string) bool { // If s is a big endian heap.
	b := make([]byte, (len(s) + 1))
	copy(b[1:], []byte(s))
	length := len(s)
	if length%2 == 0 { // Len(s) is even number. S contains a single-branch node.
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
	var s string
	fmt.Scan(&s)
	if !IsHeap(s) {
		res := HeapSort(s)
		fmt.Println(res)
	} else {
		fmt.Println(s, "is already a heap.")
	}
}