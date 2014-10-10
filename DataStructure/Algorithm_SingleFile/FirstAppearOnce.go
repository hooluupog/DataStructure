// 找出字符串中第一个仅出现一次的字符

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	table := make([]int, 256)
	for _, v := range s {
		table[v]++
	}
	for _, v := range s {
		if table[v] == 1 {
			fmt.Println(string(v))
			return
		}
	}
}
