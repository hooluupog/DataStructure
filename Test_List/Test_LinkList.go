package main

import (
	"container/list"
	"fmt"
	"sort"
)

func ReverseBetween(l *list.List, m, n int) {
	if m < 1 || n > l.Len() {
		panic(fmt.Sprintf("Index out of bounds.Wrong range (%d,%d).", m, n))
	}
	cnt := n - m
	start := l.Front()
	for i := 1; i < m; i++ { //start points to the first node to be reversed.
		start = start.Next()
	}
	end := start
	for cnt > 0 { // Insert nodes from m to n into list.
		curr := end.Next()
		l.MoveBefore(curr, start)
		start = curr
		cnt--
	}
}

func toSlice(L *list.List) []float64 {
	var res []float64
	for e := L.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(float64))
	}
	return res
}

func main() {
	l := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 4, 0, 5.0, 2, 8, 9}
	L := list.New()
	L.PushBack(float64(0))
	e := L.Front()
	for _, v := range l[0 : len(l)/2] {
		L.PushBack(v)
	}
	for _, v := range l[len(l)/2 : len(l)] {
		L.PushFront(v)
	}
	fmt.Println(toSlice(L))
	ReverseBetween(L, 3, 8)
	fmt.Println(toSlice(L))
	L.Remove(e)
	fmt.Println(toSlice(L))
	fmt.Printf("first = %v last = %v length = %v \n", L.Front().Value, L.Back().Value, L.Len())
	var sl []float64
	for p := L.Front(); p != nil; p = p.Next() {
		if p.Value.(float64) < float64(7) {
			sl = append(sl, p.Value.(float64))
		}
	}
	fmt.Println(sl)
	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	fmt.Println(sl)
	ReverseBetween(L, 1, L.Len())
	fmt.Println(toSlice(L))
	ReverseBetween(L, 0, L.Len()) // Expected error.
}
