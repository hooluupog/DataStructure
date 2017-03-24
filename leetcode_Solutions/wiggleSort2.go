package main

import (
	"fmt"
	"math/rand"
)

// 中位数将数组划分为两半，一部分大于另一部分；
// 再通过虚拟下标索引方法，重新映射数组的下标；
// 偶数位比中位数小的交换到后面，k记录较大的元素最新位置；
// 奇数位比中位数大的交换到前面，j记录较小的元素最新位置.

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
func vIndex(i, length int) int { // Calculate virtual index.
	return (2*i + 1) % (length | 1)
}
func qselect(nums []int, low, high int) int {
	mid := len(nums) >> 1
	for {
		i := low
		pivot := low + rand.Int()%(high-low+1)
		nums[pivot], nums[low] = nums[low], nums[pivot]
		for j := low + 1; j <= high; j++ {
			if nums[j] < nums[low] {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i], nums[low] = nums[low], nums[i]
		if i < mid {
			low, high = i+1, high
		} else if i > mid {
			low, high = low, i-1
		} else {
			return nums[i]
		}
	}
}

func main() {
	//nums := []int{6, 4, 6, 5, 7, 6, 8, 9}
	//nums := []int{1, 3, 2, 2, 3, 1}
	//nums := []int{4, 5, 5, 5, 5, 6, 6, 6}
	nums := []int{1, 1, 2, 1, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}
