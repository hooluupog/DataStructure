// 在数组中，数字减去它右边的数字得到一个数对之差。
// 求所有数对之差的最大值。例如在数组{2, 4, 1, 16, 7, 5, 11, 9}中，
// 数对之差的最大值是11，是16减去5的结果

package main

import "fmt"

func GreatestSum(a []int) int64 {
	// 从后向前扫描数组，min记录最小的记录值
	// sum = Max(sum,a[i] - min)
	// min = Min(min,a[i])
	sum := int64(a[len(a)-2]) - int64(a[len(a)-1])
	min := a[len(a)-1]
	for i := len(a) - 2; i >= 0; i-- {
		if sum < int64(a[i])-int64(min) {
			sum = int64(a[i]) - int64(min)
		}
		if a[i] < min {
			min = a[i]
		}
	}
	return sum
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
