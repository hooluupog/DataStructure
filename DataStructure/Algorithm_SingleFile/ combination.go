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

func combination(num []int, n int, result []int) {
	for i := n; i < len(num); i++ {
		// use temp to save the origin sequence.
		temp := make([]int, len(num)-n)
		copy(temp, num[n:])
		num[n], num[i] = num[i], num[n]
		result[n] = num[n]
		if n == len(num)-1 {
			fmt.Println(result)
		}
		combination(num, n+1, result)
		//recover sequence.
		copy(num[n:], temp)
	}
}
func main() {
	var n int
	var num, result []int
	fmt.Scanln(&n)
	num = make([]int, n)
	result = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num[i])
	}
	combination(num, 0, result)
}
