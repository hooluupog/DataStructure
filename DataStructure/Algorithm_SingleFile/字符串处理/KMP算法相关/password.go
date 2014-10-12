/* 
* codeforces-127D
* Find the longest substring p which is prefix,suffix and infix of string s.
* sample
* input
* fixprefixsuffix
* output
* fix
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func password(s string) {
	if len(s) <= 2 {
		fmt.Println("Just a legend")
		return
	}
	next := make([]int, len(s)+1)
	count := make([]int, len(s)+1)
	next[0] = -1
	for i, j := 0, -1; i < len(s); {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
			count[j]++
		} else {
			j = next[j]
		}
	}
	// 由kmp算法可知，s[0 : next[(len(s)]]必为最长的前缀和后缀子串
	// 接下来判断其是否为中缀。分两种情况：
	// 1)该子串出现的次数(匹配次数)大于1，说明其必为中缀；
	// 2)该子串出现次数不大于1，则考察next[next[len(s)]]，如果大于0，
	// 则s[0:next[next[len(s)]]]必为最长中缀,同时也是前缀和后缀。
	if count[next[len(s)]] > 1 {
		s = s[:(next[len(s)])]
	} else {
		s = s[:(next[next[len(s)]])]
	}
	if s != "" {
		fmt.Println(s)
	} else {
		fmt.Println("Just a legend")
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.TrimRight(s, "\n\r")
	s = strings.TrimRight(s, "\n")
	password(s)
}
