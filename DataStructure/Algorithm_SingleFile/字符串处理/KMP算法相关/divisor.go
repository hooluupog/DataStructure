/*
* find the common divisor of string s1 and s2. The divisor is prefix of
* a string or itself.
* sample input
* abcdabcd
* abcdabcdabcdabcd
* sample output
* 2
 */
package main

import "fmt"

func next(p string) []int {
	next := make([]int, len(p)+1)
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
	return next
}

func divisor(s1, s2 string) {
	var count int
	next1 := next(s1)
	next2 := next(s2)
	minFix1 := len(s1) - next1[len(s1)]
	minFix2 := len(s2) - next2[len(s2)]
	switch {
	case len(s1)%minFix1 == 0 && len(s2)%minFix2 == 0:
		// Both s1 and s2 contain the shortest divisor.
		if s1[:minFix1] == s2[:minFix2] {
			for i := minFix1; i <= len(s1) && i <= len(s2); i = minFix1 + i {
				if (len(s1)%i + len(s2)%i) == 0 {
					count++
				}
			}
		}
	case len(s1)%minFix1 != 0 && len(s2)%minFix2 == 0:
		// s1 has no divisor except itself.
		if len(s1) < len(s2) && len(s1)%minFix2 == 0 && s1 == s2[:len(s1)] {
			count = 1
		}
	case len(s1)%minFix1 == 0 && len(s2)%minFix2 != 0:
		// s2 has no divisor except itself.
		if len(s2) < len(s1) && len(s2)%minFix1 == 0 && s2 == s1[:len(s2)] {
			count = 1
		}
	case len(s1)%minFix1 != 0 && len(s2)%minFix2 != 0:
		//Both s1 and s2 have none divisors.
		if s1 == s2 {
			count = 1
		}
	}
	fmt.Println(count)
}
func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	divisor(s1, s2)
}
