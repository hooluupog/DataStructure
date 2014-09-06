/*
// kmp算法的应用
// HDU--2203题 亲和串
// 亲和串的定义是这样的：给定两个字符串s1和s2,如果能通过s1循环移位，
// 使s2包含在s1中，那么我们就说s2 是s1的亲和串。
// Sample Input
// AABCD
// CDAA
// ASD
// ASDF
//
// Sample Output
// yes
// no
//
*/
package main

import (
	"fmt"
	"io"
	"kmp"
	"log"
)

func AmeicableString(s string, p string) int {
	var res []int
	next := make([]int, len(p))
	next = kmp.Get_next(p)
	if len(s) < len(p) {
		return 0
	}
	for i := 0; i < len(s); i++ {
		res = kmp.Kmp(s[i:]+s[:i], p, next)
		if len(res) != 0 {
			return 1
		}
	}
	return 0
}

func main() {
	var s, p string
	for {
		_, err := fmt.Scan(&s, &p)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		result := AmeicableString(s, p)
		if result == 1 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
