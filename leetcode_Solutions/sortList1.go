package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// sort linkList with O(nlgn).
func sortList(head *ListNode) *ListNode {
	if head != nil {
		rear := head
		for rear.Next != nil {
			rear = rear.Next
		}
		qsort(head, rear)
	}
	return head
}

func qsort(head, rear *ListNode) {
	if head != rear {
		prePivot := partition(head, rear)
		if prePivot.Next != head {
			qsort(head, prePivot)
		}
		if prePivot.Next != rear {
			qsort(prePivot.Next.Next, rear)
		}
	}
}

func partition(head, rear *ListNode) *ListNode { //关键在返回枢轴结点的前驱结点
	i := head
	pre := &ListNode{}
	pre.Next = i
	for j := head.Next; j != nil; j = j.Next {
		if j.Val < head.Val {
			pre = i
			i = i.Next
			i.Val, j.Val = j.Val, i.Val
		}
	}
	head.Val, i.Val = i.Val, head.Val
	return pre
}

func printS(res *ListNode) {
	for res != nil {
		fmt.Printf("%v", res.Val)
		res = res.Next
	}
	fmt.Println()
}

func main() {
	var s string
	fmt.Scan(&s)
	head := &ListNode{}
	q := head
	for _, v := range s {
		a, _ := strconv.Atoi(string(v))
		list := &ListNode{a, nil}
		q.Next = list
		q = q.Next
	}
	res := sortList(head.Next)
	printS(res)
}
