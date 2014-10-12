// 已知有序数组a[N]，从中间某个位置k（k未知，k=-1表示整个数组有序）分开，
// 然后将前后两部分互换，得到新的数组，在该新数组的查找最小元素。
// 如：a[]={1,2,5,7,9,10,15}，从k=4分开，得到新数组a={9,10,15, 1,2,5,7}。

package main

import "fmt"

func RotateArrayMin(a []int) int {
	low, high := 0, len(a)-1
	mid := 0
	min := a[0]
	for low <= high {
		mid = (low + high) >> 1
		if a[mid] >= a[low] { // [low,mid]有序,a[low]为区间最小值
			if a[low] < min { // min和a[low]比较大小
				min = a[low]
			}
			low = mid + 1 // 查找另一部分
		} else { // [mid,high]有序,a[mid]为区间最小值
			if a[mid] < min { // min和a[mid]比较大小
				min = a[mid]
			}
			high = mid - 1 // 查找另一部分
		}
	}
	return min
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(RotateArrayMin(a))
}
