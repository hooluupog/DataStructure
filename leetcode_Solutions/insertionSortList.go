package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, s := head.Next, head
	for p != nil {
		if s.Val <= p.Val && s != p {
			s = s.Next
		} else {
			q := head
			for ; q != p && q.Val <= p.Val; q = q.Next {
			}
			if q != s {
				s.Next = p.Next
				p.Val, q.Val = q.Val, p.Val
				p.Next = q.Next
				q.Next = p
				p = s
			} else {
				p.Val, q.Val = q.Val, p.Val
				s = s.Next
			}
		}
		p = p.Next
	}
	return head
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
	res := insertionSortList(head.Next)
	printS(res)
}
