// 求数组的逆序数对总和
// 例如{4,1,3,2},逆序数对为：{4,1},{4,3},{4,2},{3,2}
// 则逆序数对总数为4
// 利用归并排序来求解此类问题：
// 思路：
// 可以利用树状数组，比如2,1,4,3：
// 
//                                 ［2,1,4,3］                                                                level 1
//                                 /        \
//                               [2,1]      [4,3]              level 2
//                              /     \     /   \
//                            [2]    [1]  [4]   [3]            level 3
// 
// 对于序列[l, r]，逆序数对分为以下三种情况：
//  1)两个数都在[l,mid]内;
//  2)两个数字都在[mid+1,r]内;
//  3)一个数在[l,mid]，一个数在[mid+1,r]内;
//  [l,r]的逆序数 = [l,mid]的逆序数 + [mid+1,r]的逆序数 + [l,mid],[mid+1,r]之间的逆序数.
//  递归向下，当区间内只有2个数时，则只会出现第3)种情况，即逆序数对分布在两个子区间内.             
// 说明1)和2)两种情况的逆序数可以由3)得出，即一个区间内的逆序数就是其对 应的下层的子区间之间的逆序数。
// 例如[2,1,4,3]的逆序数就是它的下层：[2,1],[4,3]子区间之间,[2],[1]之间，[4],[3]之间的所有逆序数。
// 累加每层的区间之间的逆序数即可得到整个序列的逆序数总和。
// 利用归并排序，在每次归并前，计算每一层的序列之间的逆序数之和，最终将所有层的和加起来即为总的逆序数和。



package main

import "fmt"

func Merge(aux []int, a []int, low int, mid int, high int) {
	for i, j, k := low, mid+1, low; k <= high; k++ {
		if j > high || i <= mid && aux[i] <= aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

func LowerBound(a []int, low, high int, value int) int {
	// 查找升序序列中value的插入位置,即比value小最后一个a[i]元素的位置
	l, h := low, high
	pos := low - 1
	for l <= h {
		mid := (l + h) >> 1
		if value > a[mid] {
			l = mid + 1
			pos = mid
		} else {
			h = mid - 1
		}
	}
	return pos
}

func InversionNumber(aux []int, a []int, low int, high int) int { // 统计逆序数
	if low >= high {
		return 0
	}
	mid := (low + high) >> 1
	count := InversionNumber(a, aux, low, mid) + InversionNumber(a, aux, mid+1, high)
	for i, j := low, mid; i <= mid; i++ { // [low,mid]和[mid+1,high]之间的逆序数
		j = LowerBound(aux, j+1, high, aux[i])
		count = count + j - mid
	}
	if aux[mid] <= aux[mid+1] { // 已经有序，直接复制到a当中
		copy(a[low:high+1], aux[low:high+1])
	} else {
		Merge(aux, a, low, mid, high)
	}
	return count
}

func main() {
	a := []int{8,7,1,4,2,6,3,5}
	aux := make([]int, len(a))
	copy(aux, a)
	sum := InversionNumber(aux, a, 0, len(a)-1)
	fmt.Println(sum)
}
