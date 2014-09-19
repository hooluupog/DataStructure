/*快速排序，使用随机数产生枢轴*/
package main

import (
	"bufio"
	"fmt"
	"io"
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

// 划分算法
func Partition(a []int, low int, high int) int {
	rIndex := low + rand.Int()%(high-low+1)
	a[rIndex], a[low] = a[low], a[rIndex] // 将枢值交换到第一个元素
	pivotkey := a[low]                    // 置当前表中第一个元素为枢轴值
	i := low
	for j := low + 1; j <= high; j++ { // 从第二个元素开始找小于基准的元素
		if a[j] < pivotkey { // 找到和交换到前面
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i], a[low] = a[low], a[i] // 将基准元素插入到最终位置
	return i
}

func main() {
	var a int
	var l []int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscan(r, &a)
		if err == io.EOF {
			break
		}
		l = append(l, a)
	}
	Qsort(l, 0, len(l)-1)
	fmt.Fprintln(w, l)
	w.Flush()
}
