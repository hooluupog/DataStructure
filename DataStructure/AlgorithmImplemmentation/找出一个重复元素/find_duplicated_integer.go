/**
*问题描述:
*取值为[1，n-1]含n个元素的整数数组至少存在一个重复数，O（n）时间内找出其中任意一个重复数。
*如a[]={1,2,2,4,5,4}，则2和4均是重复元素。
*算法:使用hash算法
*如果数组元素为signed int，可采用该方案：将数组元素值作为索引，
*对于元素a[i]>0，
*      如果a[a[i]]大于0,则设置a[a[i]]=- a[a[i]]；
*      如果a[a[i]]小于0，则a[i]是一个重复数，直接输出.
*对于元素a[i]<0,
*      如果a[-a[i]]大于0,则设置a[-a[i]]=- a[-a[i]]；
*      如果a[-a[i]]小于0，则-a[i]是一个重复数，直接输出.
*最后还原a中各个被修改的元素。
 */

package main

import "fmt"

func find_duplicated_integer(a []int) (value int) {
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			if a[a[i]] > 0 {
				a[a[i]] = -a[a[i]]
			} else {
				value = a[i]
				return
			}
		} else {
			if a[-a[i]] > 0 {
				a[-a[i]] = -a[-a[i]]
			} else {
				value = -a[i]
				return
			}
		}
	}
	fmt.Println(value)
	return
}

func reset_array(a []int) {
	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			a[i] = -a[i]
		}
	}
}

func main() {
	a := []int{1, 2, 9, 5, 4, 6, 7, 3, 6, 4}
	result := find_duplicated_integer(a)
	fmt.Println(result)
	reset_array(a)
}
