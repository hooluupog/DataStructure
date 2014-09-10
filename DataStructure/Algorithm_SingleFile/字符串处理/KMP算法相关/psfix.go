/*
* codeforces-432D
* For any prefix of string s which matches a suffix of string s, 
* print the number of times it occurs in string s as a substring.
* sample input
* ABACABA
* output
* 3
* 1 4
* 3 2
* 7 1
* algorithm: kmp,DP
*/

package main

import "fmt"

type Result struct {
	Length int
	Count  int
}

func psfix(p string) {
	var res []Result
	next := make([]int, len(p)+1)
	dp := make([]int, len(p)+1)
	next[0] = -1
	for i, j := 0, -1; i < len(p); {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	for i := len(p); i > 0; i-- {
		// dp[i] represents the count of substring p[0:i] apprearing from 0 to i.
		// its initialized value is 1.
		dp[i]++
		dp[next[i]] = dp[next[i]] + dp[i]
	}
	res = append(res, Result{len(p), 1})
	for i := next[len(p)]; i != 0; i = next[i] {
		res = append(res, Result{i, dp[i]})
	}
	fmt.Println(len(res))
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d %d\n", res[i].Length, res[i].Count)
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	psfix(s)
}
