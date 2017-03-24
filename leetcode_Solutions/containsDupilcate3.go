package main

import (
	"fmt"
	"math"
)

// 设置K个桶，依次为：[0,t],[t+1,2t],...,[(k-1)t+1,kt]
// 满足 |a1 - a2| <= t的元素，只能为两种情况：
// 1）a1,a2都被装在一个桶内（即同一个区间内）;
// 2）a1,a2在相邻的两个桶中。
// 算法时间复杂度为O(N).如果采用Set,TreeSet,TreeMap（内部为红黑树实现）
// 则算法时间复杂度为O(Nlgk).哈希表的查找时间为O(1),后者为O(lgk).

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}
	cd := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		bkt := nums[i] / (t + 1)            // bucket
		for j := bkt - 1; j <= bkt+1; j++ { // appearance in bkt or in bkt's neighbor.
			if _, ok := cd[j]; ok && math.Abs(float64(nums[i]-cd[j])) <= float64(t) {
				return true
			}
		}
		cd[bkt] = nums[i]
		if i >= k {
			delete(cd, nums[i-k]/(t+1))
		}
	}
	return false
}

func main() {
	a := []int{8, 2, 1, 3, 0}
	fmt.Println(containsNearbyAlmostDuplicate(a, 3, 5))
	b := []int{-3, 3}
	fmt.Println(containsNearbyAlmostDuplicate(b, 2, 4))
}
