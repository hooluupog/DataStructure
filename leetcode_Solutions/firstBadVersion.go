package main

import (
	"fmt"
)

const length = 2126753391

var bad [length]int

func isBadVersion(version int) bool {
	if bad[version] == 1 {
		return true
	} else {
		return false
	}
}

func findFirst(start, end int) int {
	if start == end || end < start {
		return start
	}
	mid := (start + end) >> 1
	if isBadVersion(mid) {
		return findFirst(start, mid-1)
	} else {
		return findFirst(mid+1, end)
	}
}

func firstBadVersion(n int) int {
	return findFirst(1, n)
}

func main() {
	bad[0] = -1
	for i := 1; i < length; i++ {
		if i < 1702766719 {
			bad[i] = 0
		} else {
			bad[i] = 1
		}
	}
	fmt.Println(firstBadVersion(length - 1))
}
