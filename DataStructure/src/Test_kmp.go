package main

import (
	"fmt"
	"io"
	"kmp"
	"log"
)

func main() {
	var s, p string
	//s := "acabaabaabcacaabc"
	//p := "abaabcac"
	//s := "aabacaabaabaadaabaaeaa"
	//p := "aabaa"
	//s := "ababaaca"
	//p := "baac"
	//s := "bdbcbabcbab"
	//p := "bdbcba"
	for {
		_, err := fmt.Scan(&s, &p)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		next := make([]int, len(p))
		next = kmp.Get_next(p)
		fmt.Printf("The matching postion  : %v\n", kmp.Kmp(s, p, next))
	}

}