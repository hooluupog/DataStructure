/*原始的2路归并排序，未优化*/

package main

import (
	"fmt"
	"io"
)

func MergeSort(a []int, low, high int) {
	if low < high {
		mid := (low + high) / 2   // 从中间划分两个子序列
		MergeSort(a, low, mid)    // 递归排序左子序列
		MergeSort(a, mid+1, high) // 递归排序右子序列
		Merge(a, low, mid, high)  // 归并
	}
}

//合并两个有序子序列为一个有序表
func Merge(a []int, low, mid, high int) {
	buffer := make([]int, len(a)) // 辅助数组
	copy(buffer, a)
	i, j, k := low, mid+1, low
	for ; i <= mid && j <= high; k++ {
		if buffer[i] <= buffer[j] { 
			a[k] = buffer[i] // 将较小的值复制到a中
			i++
		} else {
			a[k] = buffer[j]
			j++
		}
	} // for
	for i <= mid { // 剩余第一个表，复制
		a[k] = buffer[i]
		i++
		k++
	}
	for j <= high { // 剩余第二个表，复制
		a[k] = buffer[j]
		j++
		k++
	}
}

func main() {
	var a int
	var l []int
	for {
		_, err := fmt.Scan(&a)
		if err == io.EOF {
			break
		}
		l = append(l, a)
	}
	MergeSort(l, 0, len(l)-1)
	fmt.Println(l)
}
