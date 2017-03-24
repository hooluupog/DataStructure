package main

import (
	"fmt"
)

// 利用栈的特性，当前值大于栈顶元素则记录并出栈，否则入栈。
func nextGreaterElement(findNums []int, nums []int) []int {
	fg := make(map[int]int) // Store first greater element.
	var stack []int
	for _, v := range nums {
		for len(stack) != 0 && stack[len(stack)-1] < v { // pop
			fg[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v) // push
	}
	for i, v := range findNums {
		if m, ok := fg[v]; ok {
			findNums[i] = m
		} else {
			findNums[i] = -1
		}
	}
	return findNums
}

func main() {
	a := []int{-1, 3, -4, -2}
	b := []int{-4, -1}
	fmt.Println(nextGreaterElement(b, a))
}
