package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	for i := 1; i < len(nums); i++ { //先找出最大的
		if nums[0] < nums[i] {
			nums[0], nums[i] = nums[i], nums[0]
		}
	}
	for i := 1; i <= 2; i++ { //再依次找出次大的
		j := i
		for j < len(nums) && nums[j] >= nums[i-1] {
			j++
		}
		if j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
			for j = i + 1; j < len(nums); j++ {
				if nums[i] < nums[j] && nums[j] < nums[i-1] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
	}
	if len(nums) < 3 || nums[1] <= nums[2] {
		return nums[0]
	} else {
		return nums[2]
	}
}
func main() {
	a := []int{3, 3, 4, 3, 4, 3, 0, 3, 3}
	fmt.Println(thirdMax(a))
}
