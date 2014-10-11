package main

import "fmt"

func GreatestSum(a []int) int64 {
	// 动态规划法
	// diff[i] = max{a[0],a[1],...,a[i-1]} - a[i]
	// 则diff[i+1] = max{a[0],...,a[i]} - a[i+1]
	// ans = max{diff[0],diff[1],...,diff[n-1]}
	// 其中，diff[i]表示以a[i]为减数的所有数对之差的最大值

	if len(a) < 2 {
		return 0
	}
	max := a[0]
	ans := int64(max) - int64(a[1])
	for i := 2; i < len(a); i++ {
		if a[i-1] > max {
			max = a[i-1]
		}
		diff := int64(max) - int64(a[i])
		if diff > ans {
			ans = diff
		}
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(GreatestSum(a))
}
