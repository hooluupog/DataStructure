/**
*问题描述:
*取值为[1，n-1]含n个元素的整数数组至少存在一个重复数，O（n）时间内找出其中任意一个重复数。
*如a[]={1,2,2,4,5,4}，则2和4均是重复元素。
*算法:利用单链表存在环
*“判断单链表是否存在环”是一个非常经典的问题，同时单链表可以采用数组实现，此时每个元素值作
*为next指针指向下一个元素。该题可以转化为“已知一个单链表中存在环，找出环的入口点”。
*该题思路如下：将a[i]看做第i个元素的索引，即：a[i]->a[a[i]]->a[a[a[i]]]->a[a[a[a[i]]]]->….
*最终形成一个单链表，由于数组a中存在重复元素，则一定存在一个环，且环的入口元素即为重复元素。
*该题的关键在于，数组a的大小时n，而元素的范围是[1,n-1]，所以a[0]不会指向自己，进而不会陷入错
*误的自循环。如果元素的范围中包含0，则该题不可直接采用该方法。
 */

package main

import "fmt"

func find_duplicated_integer(a []int) (value int) {
	x, y := 0, 0
	//找出环中的相遇点
	for x, y = a[a[0]], a[0]; x != y; x, y = a[a[x]], a[y] {
	}
	//寻找环的入口点
	//x从表头开始向后移动,y从相遇点的下一个点开始向后移动(即想象将环从相遇点处
	//切开,问题转化为求两个单链表交点的问题)
	for x, y = a[0], a[y]; x != y; x, y = a[x], a[y] {
	}
	value = x
	return
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 4}
	result := find_duplicated_integer(a)
	fmt.Println(result)
}
