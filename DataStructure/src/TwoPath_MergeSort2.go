/*
* 使用直接插入排序进行优化后的2路归并排序
* 插入排序在序列长度n很小的情况下拥有较快的性能，因此可以通过
* 插入排序来优化归并排序的子序列，减少递归深度。
* 因此对长度为n的线性表进行归并排序的问题转化为对长度为n/k的k个
* 子序列进行插入排序，最后将这些有序的子序列进行归并排序。
* 那么k应该如何取值?
* 显然：k应该是使插入排序比归并排序快的最大的子表长度。
* 假设：有n/k个具有k个元素的列表，对每个列表采用直接插入排序，最后归并已有序的列表完成整个排序。
* 可以得到如下的分析：
* 对每个列表排序的最坏时间是Θ(k^2),则n/k个列表需要Θ(nk)的时间。
* 合并这些列表需要Θ(nlg(n/k))的时间。(原始归并排序中，合并n/k个规模为k的列表需要 cn/k * k = Θ(n)，
* 利用数学归纳法可证每次合并都需要Θ(n)，并且需要lg(n/k)次合并)。
* 总时间为Θ(nk+nlg(n/k)).
* 因此有,Θ(nk+nlg(n/k)) = Θ(n(k + lgn - lgk)) = Θ(nlgn) => k = Θ(lgn) 即k的阶取值不能大于Θ(lgn).
* 实际中，k应当选择使插入排序比归并排序快的最大的子表长度，即cnk + dnlg(n/k)执行时间要最小, c 和 d是正常数，
* 分别决定着实际环境中两个部分的各种执行时间(在不同的机器和平台上会有差别).当k = d/c 时，上式取得最小值。
* 所以，k的取值和n无关，而是由常数d和c的比决定。
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const INSERTSORT_THRESHOLD = 7

func MergeSort(aux []int, a []int, low, high int) {
	if low < high {
		length := high - low + 1
		// 小数组采用插入排序
		if length < INSERTSORT_THRESHOLD {
			InsertSort(a, low, high)
			return
		}
		// Recursively sort halves of a into aux.
		mid := (low + high) >> 1 // 从中间划分两个子序列
		//将a看作辅助数组，aux看作目标数组，如此反复，直到全部有序
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
