package main

import (
	"fmt"
	"math/rand"
)

func wiggleSort(nums []int) {
	median := qselect(nums, 0, len(nums)-1)
	length := len(nums)
	for i, j, k := 0, 0, length-1; i <= k; {
		iv, jv, kv := vIndex(i, length), vIndex(j, length), vIndex(k, length)
		if nums[iv] > median {
			nums[iv], nums[jv] = nums[jv], nums[iv]
			i++
			j++
		} else if nums[iv] < median {
			nums[iv], nums[kv] = nums[kv], nums[iv]
			k--
		} else {
			i++
		}
	}
}
func vIndex(i, length int) int { // virtual index.
	return (2*i + 1) % (length | 1)
}
func qselect(nums []int, low, high int) int {
	mid := len(nums) >> 1
	if low < high {
		pivot := partition(nums, low, high)
		if pivot < mid {
			return qselect(nums, pivot+1, high)
		}
		if pivot > mid {
			return qselect(nums, low, pivot-1)
		}
	}
	return nums[mid]
}
func partition(nums []int, low, high int) int {
	pivot := low + rand.Int()%(high-low+1)
	nums[pivot], nums[low] = nums[low], nums[pivot]
	i := low
	for j := low + 1; j <= high; j++ {
		if nums[j] < nums[low] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[low] = nums[low], nums[i]
	return i
}
func main() {
	//nums := []int{6, 4, 6, 5, 7, 6, 8, 9}
	//nums := []int{1, 3, 2, 2, 3, 1}
	//nums := []int{4, 5, 5, 5, 5, 6, 6, 6}
	nums := []int{1, 1, 2, 1, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}
