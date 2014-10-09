// 同OddEven.cpp,只是这个算法会改变数组中奇数和偶数的原次序位置

package main

import "fmt"

func OddEven(a []int) {
	pivotkey, low := 0, 0
	for ; low < len(a) && a[low]%2 == 0; low++ {
	}
	if low < len(a) {
		pivotkey = a[low]
		for i := low; i > 0; i-- {
			a[i] = a[i-1]
		}
		a[0] = pivotkey
	}
	odd, even := 1, 1
	for ; even < len(a); even++ {
		if a[even]%2 != 0 {
			pivotkey = a[even]
			for i := even; i > odd; i-- {
				a[i] = a[i-1]
			}
			a[odd] = pivotkey
			odd++
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	OddEven(a)
	fmt.Println(a)
}
