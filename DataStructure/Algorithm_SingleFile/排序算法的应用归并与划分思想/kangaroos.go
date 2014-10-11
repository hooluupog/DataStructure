/*
* codeforces 373c
* 有n只袋鼠，每只袋鼠有一个袋子，当一只袋鼠的袋子大小(数组的整数值)至少为另一只袋鼠的两倍大时，可以装进另一只袋鼠，
* 一只袋鼠只能装一只袋鼠，被装的袋鼠不能装其他袋鼠。被装进去的袋鼠不可见。找出一个方案使得可见的袋鼠数量最少。
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func Qsort(a []int, low int, high int) {
	if low < high {
		pivot := Partition(a, low, high)
		Qsort(a, low, pivot-1)
		Qsort(a, pivot+1, high)
	}
}

func Partition(a []int, low int, high int) int {
	pivot := low + rand.Int()%(high-low+1)
	a[pivot], a[low] = a[low], a[pivot]
	i := low
	for j := low + 1; j <= high; j++ {
		if a[j] < a[low] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i], a[low] = a[low], a[i]
	return i
}

func Kangaroos(a []int) {
	n := len(a)
	visable := n
	for low, high := n/2-1, n-1; low >= 0; low-- {
		if a[low]*2 <= a[high] {
			visable--
			high--
		}
	}
	fmt.Println(visable)
}

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &a[i])
	}
	Qsort(a, 0, n-1)
	Kangaroos(a)
}
