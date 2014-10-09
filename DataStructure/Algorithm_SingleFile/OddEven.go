// 同OddEven.cpp

package main

import "fmt"

func OddEven(a []int) {
	pivotkey, low := 0, 0
	// 将第一个奇数交换到数组第一位作为枢轴
	for ; low < len(a) && a[low]%2 == 0; low++ {
	}
	if low < len(a) {
		pivotkey = a[low]
		for i := low; i > 0; i-- {
			a[i] = a[i-1]
		}
		a[0] = pivotkey
	}
	// 向后扫描，发现奇数将其移动到上一个奇数的下一位，后面的其他元素依次后移。
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
