package main

import "fmt"

func singleNonDuplicate(nums []int) int { //长度为奇数或偶数，两种情况分别处理
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		if (mid-start)%2 != 0 {
			if nums[mid] == nums[mid+1] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] == nums[mid+1] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	return nums[start]
}

func main() {
	a := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(a))
	b := []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Println(singleNonDuplicate(b))
}
