package main

import (
	"fmt"
)

//利用栈的特性

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	ng := make([]int, len(nums)) // Store next greater element.
	for v := range ng {
		ng[v] = -1
	}
	var stack []int
	max := nums[0]
	i, j := 0, 0
	stop := false
	for {
		if max < nums[i] {
			max = nums[i]
		}
		// 循环搜索的结束判断
		if j >= len(nums) && max == nums[i] { // the last time traversing.
			stop = true
		}
		// 比当前栈顶元素大的话，记录并出栈；否则入栈。
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] { // pop
			ng[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		if stop {
			break
		}
		stack = append(stack, i) // push
		i = (i + 1) % len(nums)
		j++
	}
	return ng
}

func main() {
	a := []int{2, 3, 2}
	fmt.Println(nextGreaterElements(a))
}
