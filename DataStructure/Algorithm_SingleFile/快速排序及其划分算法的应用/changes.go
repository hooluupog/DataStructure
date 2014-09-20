/*
* codeforces 262B
* 数组大小为n,数组序列初始状态是升序的，进行k次转换(转换：对当前的元素值乘以-1),一个元素可以进行多次转换。
* 求经过k次转换后，数组元素之和的最大值。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Income(a []int, k int) {
	n := len(a)
	sum := 0
	if k < n && a[k-1] < 0 { // 多于k个负数
		for i := 0; i < k; i++ {
			a[i] *= -1
		}
	} else { // 全部转变为正数
		changes := 0 //转换的次数
		for i := 0; i < k && i < n; i++ {
			if a[i] < 0 { // 负数变正
				a[i] *= -1
				changes++
			} else { // 已经没有负数，跳出循环
				break
			}
		}
		sort.IntSlice(a).Sort() // 对正数序列重新排序.
		k = k - changes         // 剩余的转变次数.
		if k%2 != 0 {           // 转变次数为奇数次
			a[0] *= -1 // 将最小的正数转变为负数
		}
	}
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}

func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Scanln(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	Income(a, k)
}
