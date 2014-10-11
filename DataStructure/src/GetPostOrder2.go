package main

import (
	"fmt"
	"io"
	"runtime"
)

var pre,in []int

func postOrder(pstart, istart int, n int, isRoot bool) {
	if n < 1 { return }
	i := 0
	for ; pre[pstart] != in[istart+i]; i++ {}
	postOrder(pstart+1, istart, i, false)
	postOrder(pstart+1+i, istart+i+1, n-i-1, false)
	if isRoot {
		fmt.Printf("%d", pre[pstart])
	} else {
		fmt.Printf("%d ", pre[pstart])
	}
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		pre = make([]int, n)
		in = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&pre[i])
		}
		for i := 0; i < n; i++ {
			fmt.Scan(&in[i])
		}
		postOrder(0, 0, n, true)
		fmt.Printf("\n")
	}
}
