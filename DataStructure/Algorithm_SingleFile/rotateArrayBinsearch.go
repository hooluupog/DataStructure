// 已知有序数组a[N]，从中间某个位置k（k未知，k=-1表示整个数组有序）分开，
// 然后将前后两部分互换，得到新的数组，在该新数组的查找元素e。
// 如：a[]={1,2,5,7,9,10,15}，从k=4分开，得到新数组a={9,10,15, 1,2,5,7}。

package main

import "fmt"

func RotateBinSearch(a []int, e int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := (low + high) >> 1
		if a[mid] == e {
			return mid + 1
		}
		if a[mid] >= a[low] { // [low,mid]有序
			if e < a[mid] && e >= a[low] { // e属于[low,mid-1]有序区间
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // [mid+1,high]有序
			if e > a[mid] && e <= a[high] { // e属于[mid+1,high]有序区间
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return 0
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(RotateBinSearch(a, e))
}
