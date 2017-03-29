// 给一个整数，求所有bit位上1的个数。
// 算法：迭代采用 n & (n - 1)将n的最右的1bit位清除 , 直到所有bit位为0.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func setBits(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt
}

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println(setBits(n))
}
