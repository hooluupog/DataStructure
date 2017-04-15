package main

import (
	"bufio"
	"fmt"
	slist "list/intrusive"
	"os"
	"sort"
	"strconv"
	"time"
)

// Here is how to use intrusive linked list:
// 1) Define Data struct embedded Elem;
// 2) Implement item interface.

type E struct {
	val interface{}
	slist.Elem
}

func (e *E) Get() *slist.Elem { return &e.Elem }

//////////////////////////////////////////////////////////////

func toSlice(L *slist.List) []float64 {
	var res []float64
	for e := L.First(); e != nil; e = e.Get().Next() {
		res = append(res, e.(*E).val.(float64))
	}
	return res
}

func removeTest() {
	var l []*E
	L := slist.New()
	e := &E{val: float64(0)}
	L.Add(e)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		f, _ := strconv.ParseFloat(scanner.Text(), 64)
		l = append(l, &E{val: f})
	}
	for _, v := range l[0 : len(l)/2] {
		L.Add(v)
	}
	for _, v := range l[len(l)/2 : len(l)] {
		L.AddFirst(v)
	}
	start := time.Now()
	L.Remove(e)
	i := len(l) / 2
	j := i + 1
	for L.Len() != 0 {
		if i >= 0 {
			L.Remove(l[i])
			i--
		}
		if j < len(l) {
			L.Remove(l[j])
			j++
		}
	}
	duration := time.Since(start)
	fmt.Printf(" %vms\n", duration.Seconds()*1000)
}

func main() {
	removeTest()
	l := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 4, 0, 5.0, 2, 8, 9}
	L := slist.New()
	e := &E{val: float64(0)}
	L.Add(e)
	for _, v := range l[0 : len(l)/2] {
		L.Add(&E{val: float64(v)})
	}
	for _, v := range l[len(l)/2 : len(l)] {
		L.AddFirst(&E{val: float64(v)})
	}
	fmt.Println(toSlice(L))
	L.ReverseBetween(3, 8)
	fmt.Println(toSlice(L))
	L.Remove(e)
	fmt.Println(toSlice(L))
	fmt.Printf("first = %v last = %v length = %v \n", L.First().(*E).val, L.Last().(*E).val, L.Len())
	nl := toSlice(L)
	var sl []float64
	for _, v := range nl {
		if v < float64(7) {
			sl = append(sl, v)
		}
	}
	fmt.Println(sl)
	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	fmt.Println(sl)
	L.ReverseBetween(1, L.Len())
	fmt.Println(toSlice(L))
	L.ReverseBetween(0, L.Len()) // Expected error.
}
