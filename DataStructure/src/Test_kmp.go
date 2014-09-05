package main

import (
	"fmt"
	"kmp"
)

func main() {
	sstring := "acabaabaabcacaabc"
	pstring := "abaabcac"
	//sstring := "aabacaabaabaadaabaaeaa"
	//pstring := "aabaa"
	next := make([]int, len(pstring))
	next = kmp.Get_next(pstring)
	fmt.Printf("The matching postion  : %v\n", kmp.Kmp(sstring, pstring, next))
}