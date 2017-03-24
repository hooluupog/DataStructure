package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var work []*ListNode
	for head != nil {
		work = append(work, head)
		head = head.Next
	}
	qsort(work, 0, len(work)-1)
	res := work[0]
	q := res
	for i := 1; i < len(work); i++ {
		q.Next = work[i]
		q = q.Next
	}
	q.Next = nil
	return res
}

func qsort(work []*ListNode, low, high int) {
	if low < high {
		pivot := partition(work, low, high)
		qsort(work, low, pivot-1)
		qsort(work, pivot+1, high)
	}
}
func partition(work []*ListNode, low, high int) int {
	pivot := low + rand.Int()%(high-low+1)
	work[low], work[pivot] = work[pivot], work[low]
	i := low
	for j := low + 1; j <= high; j++ {
		if work[j].Val < work[low].Val {
			i++
			work[j], work[i] = work[i], work[j]
		}
	}
	work[low], work[i] = work[i], work[low]
	return i
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
