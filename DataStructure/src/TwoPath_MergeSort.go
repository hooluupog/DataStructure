/*
* 标准的2路归并排序
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func MergeSort(aux []int, a []int, low, high int) {
	if low < high {
		// Recursively sort halves of a into aux.
		mid := (low + high) >> 1 // 从中间划分两个子序列
		// 将a看作辅助数组，aux看作目标数组，如此反复，直到全部有序
		MergeSort(a, aux, low, mid)    // 递归排序左子序列
		MergeSort(a, aux, mid+1, high) // 递归排序右子序列
		// 已经有序的子序列直接复制到a,避免重复排序已经有序的序列.
		if aux[mid] <= aux[mid+1] {
			copy(a[low:high+1], aux[low:high+1])
			return
		}
		Merge(aux, a, low, mid, high) // 归并
	}
}

func Merge(aux []int, a []int, low, mid, high int) {
	for i, j, k := low, mid+1, low; k <= high; k++ {
		if j > high || i <= mid && aux[i] <= aux[j] {
			a[k] = aux[i] // 将较小的值复制到a中
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	} // for
}

func InsertSort(a []int, low, high int) {
	for i := low + 1; i <= high; i++ { // 依次将a[low+1]~a[high]插入到前面已排序序列
		temp := a[i]     // 暂存a[i]
		l, h := low, i-1 // 设置折半查找范围
		for l <= h {     // 开始折半查找(升序)
			mid := (l + h) / 2 // 取中间点
			if a[mid] > temp { // 查找左半子表
				h = mid - 1
			} else { // 查找右半子表
				l = mid + 1
			}
		}
		for j := i - 1; j >= h+1; j-- { // 统一后移元素，空出插入位置
			a[j+1] = a[j]
		}
		a[h+1] = temp // 插入操作
	}
}

func main() {
	var a int
	var l []int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscan(r, &a)
		if err == io.EOF {
			break
		}
		l = append(l, a)
	}
	aux := make([]int, len(l)) // 辅助数组
	copy(aux, l)
	MergeSort(aux, l, 0, len(l)-1)
	fmt.Fprintln(w, l)
	w.Flush()
}
