/*给定m个串，求长度大等于3且在每个串中都出现的最长子串，如果两个子串
 * 长度一样要求输出字典序小的一个答案。如果没有则输出
 * “no significant commonalities".
 */
package main

import (
	"fmt"
	"kmp"
)

const SIZE = 100

var (
	answer string
	patern string
	str    = make([]string, SIZE)
)

func main() {
	var n, m, j int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scan(&str[i])
		}
		answer = ""
		flag := 0
		sLen := len(str[0])
		for pLen := sLen; pLen >= 3; pLen-- {
			for i := 0; i <= sLen-pLen; i++ {
				patern = ""
				for j = i; j < i+pLen; j++ {
					patern += string(str[0][j])
				}
				for j = 1; j < m; j++ {
					next := make([]int, len(patern)+1)
					next = kmp.Get_next(patern)
					if kmp.Kmp(str[j], patern, next) == -1 {
						break
					}
				}
				if j == m {
					flag = 1
					if len(answer) == len(patern) && answer > patern {
						answer = patern
					}
					if len(answer) < len(patern) {
						answer = patern
					}
				}
			}
			if flag == 1 {
				break
			}
		}
		if answer == "" {
			fmt.Println("no significant commonalities")
		} else {
			fmt.Println(answer)
		}
	}
}
