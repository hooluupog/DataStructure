// sample input
// 3
// 1 2 3 
//
//sample output
// 1 2 3
// 1 3 2
// 2 1 3
// 2 3 1
// 3 1 2
// 3 2 1

package main

import "fmt"

func isSwap(num []int, start, end int) bool {
	for i := start; i < end; i++ {
		if num[i] == num[end] {
			return false
		}
	}
	return true
}

func combination(num []int, n int) {
	if n == len(num)-1 {
		fmt.Println(num)
		return
	}
	for i := n; i < len(num); i++ {
		temp := make([]int, len(num)-n)
		if n == i {
			combination(num, n+1)
		} else {
			// do not swap same elements
			if !isSwap(num, n, i) {
				continue
			} else {
				// use temp to save the origin sequence.
				copy(temp, num[n:])
				num[n], num[i] = num[i], num[n]
				combination(num, n+1)
				// recover sequence.
				copy(num[n:], temp)
			}
		}
	}
}

func main() {
	var n int
	var num []int
	fmt.Scanln(&n)
	num = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num[i])
	}
	combination(num, 0)
}
