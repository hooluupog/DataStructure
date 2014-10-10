// 输入一个整形数组，数组里有正数也有负数。数组中连续的一个或多个整数
// 组成一个子数组，每个子数组都有一个和。求所有子数组的和的最大值。
// 要求时间复杂度为O(n)。
// 例如输入的数组为1, -2, 3, 10, -4, 7, 2, -5，
// 和最大的子数组为3, 10, -4, 7, 2，因此输出为该子数组的和18。

package main

import "fmt"

func main() {
	var a, n, max int
	var sum, temp int64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if i == 0 {
			max = a
		} else {
			if a > max {
				max = a
			}
		}
		//和为负数则清零，和为正数，与上次的和比较，取大的
		temp += int64(a)
		if temp < 0 {
			temp = 0
		}
		if temp > sum {
			sum = temp
		}
	}
	// all data are negative. The max element is sum.
	if sum == 0 {
		sum = int64(max)
	}
	fmt.Println(sum)
}
