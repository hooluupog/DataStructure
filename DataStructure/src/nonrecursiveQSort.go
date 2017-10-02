package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const STACK_SIZE = 1000 // 栈最大深度

func partition(list []int, low, high int) int {
	mid := low + (high-low)/2
	list[low], list[mid] = list[mid], list[low] // 枢轴元素交换到首位
	j := low + 1                                // j指向第一个大于枢轴元素的位置
	for i := low; i <= high; i++ {
		if list[low] > list[i] { // 比枢轴元素小的元素交换到前面
			list[i], list[j] = list[j], list[i]
			j++
		}
	}
	list[low], list[j-1] = list[j-1], list[low] // 枢轴元素归位
	return j - 1
}

// 快速排序非递归算法实现
func nonrecursiveQSort(list []int) {
	var stack [STACK_SIZE]int
	top := -1 //指向栈顶元素
	top++
	stack[top] = 0
	top++
	stack[top] = len(list) - 1
	for top > -1 { // 栈不空
		low, high := stack[top-1], stack[top] // 记录待排子序列起、始位置
		top -= 2                              // 出栈
		pivot := partition(list, low, high)   // 一趟划分，返回枢轴下标索引
		if pivot-1 > low {                    // 子序列长度大于1，入栈
			top++
			stack[top] = low
			top++
			stack[top] = pivot - 1
		}
		if pivot+1 < high { // 子序列长度大于1，入栈
			top++
			stack[top] = pivot + 1
			top++
			stack[top] = high
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	var a []int
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		a = append(a, v)
	}
	nonrecursiveQSort(a)
	for i, v := range a {
		if i == 0 {
			fmt.Fprint(w, v)
		} else {
			fmt.Fprint(w, " ", v)
		}
	}
	w.Flush()
	fmt.Println()
}
