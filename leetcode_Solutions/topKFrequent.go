package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key   int
	value int
}
type pairList []pair

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}
	pl := make(pairList, len(cnt))
	i := 0
	for k, v := range cnt {
		pl[i] = pair{k, v}
		i++
	}
	sort.Slice(pl, func(i, j int) bool { return pl[i].value > pl[j].value })
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, pl[i].key)
	}
	return res
}

func main() {
	a := []int{2, 3, 1, 2, 1, 1}
	res := topKFrequent(a, 2)
	fmt.Println(res)
}
