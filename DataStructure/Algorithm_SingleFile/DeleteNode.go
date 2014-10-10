// O(1) 时间内删除链表结点

package main

import "fmt"

type LinkList struct {
	Data int
	Next *LinkList
}

func CreateList(L *LinkList, n int) {
	p := L
	s := L
	for i := 0; i < n; i++ {
		s = new(LinkList)
		fmt.Scan(&s.Data)
		p.Next = s
		p = s
	}
}

func PrintList(L LinkList) {
	for p := L.Next; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Data)
	}
	fmt.Println()
}

func DeleteNode(L *LinkList, p *LinkList) {
	// Delete P node within O(1).
	if p.Next != nil {
		p.Data = p.Next.Data
		p.Next = p.Next.Next
	} else {
		// p is the last node.Find p's pre node.
		s := L
		for ; s.Next != p; s = s.Next {
		}
		s.Next = nil
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	if k < 1 || k > n {
		fmt.Println("illegal delete position")
		return
	}
	L := new(LinkList)
	CreateList(L, n)
	q := L
	for i, p := 0, L; p != nil && i <= k; p = p.Next {
		i++
		q = p
	}
	PrintList(*L)
	DeleteNode(L, q)
	PrintList(*L)
}
