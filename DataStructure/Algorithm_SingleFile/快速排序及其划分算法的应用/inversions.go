/* 归并思想的应用 */
/* codeforces 414C - Mashmokh and Reverse Operation */

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const N = 20

var (
	a   []int64 = make([]int64, (1<<N)+1)
	aux []int64 = make([]int64, (1<<N)+1)
	sum [N + 1][2]int64
)

func Reverse(a []int64, low int64, high int64) {
	for i := low; i <= (low+high)/2; i++ {
		a[i], a[high-i+low] = a[high-i+low], a[i]
	}
}

func LowerBound(a []int64, low, high int64, value int64) int64 {
	l, h := low, high
	pos := low - 1
	for l <= h {
		mid := (l + h) >> 1
		if value > a[mid] {
			l = mid + 1
			pos = mid
		} else {
			h = mid - 1
		}
	}
	return pos
}

func Merge(aux []int64, a []int64, low, mid, high int64) {
	for i, j, k := low, mid+1, low; k <= high; k++ {
		if j > high || i <= mid && aux[i] <= aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}
func Inversions(aux []int64, a []int64, low int64, high int64, deep int64) {
	if low >= high {
		return
	}
	mid := (low + high) >> 1
	Inversions(a, aux, low, mid, deep-1)
	Inversions(a, aux, mid+1, high, deep-1)
	count := int64(0)
	for i, j := low, mid; i <= mid; i++ {
		j = LowerBound(aux, j+1, high, aux[i])
		count = count + j - mid
	}
	sum[deep][0] = sum[deep][0] + count
	count = 0
	for i, j := mid+1, low-1; i <= high; i++ {
		j = LowerBound(aux, j+1, mid, aux[i])
		count = count + j - low + 1
	}
	sum[deep][1] = sum[deep][1] + count
	if aux[mid] <= aux[mid+1] {
		copy(a[low:high+1], aux[low:high+1])
		return
	}
	Merge(aux, a, low, mid, high)
}

func main() {
	var n, m, que int64
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscan(r, &n)
		if err == io.EOF {
			break
		}
		for i := 0; i < len(sum); i++ {
			sum[i][0] = int64(0)
			sum[i][1] = int64(0)
		}
		length := int64(math.Pow(2, float64(n)))
		a[0] = length
		for i := int64(1); i <= length; i++ {
			fmt.Fscan(r, &a[i])
		}
		copy(aux, a)
		Inversions(aux, a, 1, a[0], n)
		fmt.Fscan(r, &m)
		for i := int64(1); i <= m; i++ {
			fmt.Fscan(r, &que)
			ans := int64(0)
			for i := que; i >= 1; i-- {
				sum[i][0], sum[i][1] = sum[i][1], sum[i][0]
			}
			for i := int64(0); i <= n; i++ {
				ans = ans + sum[i][0]
			}
			fmt.Fprintln(w, ans)
		}
	}
	w.Flush()
}
