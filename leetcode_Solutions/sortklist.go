package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 与哈夫曼树原理相似，尽可能优先合并短的以及长度相近的
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	for length > 1 {
		j := 0
		for i := 0; i < length; {
			if i+1 == length {
				lists[j] = lists[i]
			} else {
				lists[j] = mergeTwoLists(lists[i], lists[i+1])
			}
			i += 2
			j++
		}
		length = j
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	if l1 == nil {
		l3 = l2
	}
	if l2 == nil {
		l3 = l1
	}
	if l1 != nil && l2 != nil {
		var p, q1, q2 *ListNode
		if l1.Val < l2.Val {
			l3 = l1
			q2 = l2
		} else {
			l3 = l2
			q2 = l1
		}
		p = l3
		q1 = l3.Next
		for {
			for q1 != nil && q1.Val <= q2.Val {
				p = p.Next
				q1 = q1.Next
			}
			if q1 != nil {
				temp := q1
				p.Next = q2
				p = q2
				q1 = q2.Next
				q2 = temp
			} else {
				p.Next = q2
				break
			}
		}
	}
	return l3
}
func creatList(s string) *ListNode {
	head := &ListNode{}
	q := head
	parts := strings.Split(s, ",")
	for _, v := range parts {
		a, _ := strconv.Atoi(string(v))
		list := &ListNode{a, nil}
		q.Next = list
		q = q.Next
	}
	return head.Next
}
func printS(res *ListNode) {
	for res != nil {
		fmt.Printf("%v ", res.Val)
		res = res.Next
	}
	fmt.Println()
}
func main() {
	var s string
	length := 7
	var l []*ListNode
	for i := 0; i < length; i++ {
		fmt.Scan(&s)
		l = append(l, creatList(s))
	}
	res := mergeKLists(l)
	printS(res)
}
