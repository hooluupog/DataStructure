package main

import (
	"fmt"
	"sort"
	"strings"
)

type pair struct {
	value rune
	count int
}
type pairList []pair

func frequencySort(s string) string {
	cnt := make(map[rune]int)
	for _, v := range s {
		cnt[v]++
	}
	i := 0
	pl := make(pairList, len(cnt))
	for v, c := range cnt {
		pl[i] = pair{v, c}
		i++
	}
	sort.Slice(pl, func(i, j int) bool { return pl[i].count > pl[j].count })
	res := ""
	for _, v := range pl {
		res += strings.Repeat(string(v.value), v.count)
	}
	return res
}

func main() {
	s := "都好的啦好的好的好啦"
	res := frequencySort(s)
	fmt.Println(res)
}
