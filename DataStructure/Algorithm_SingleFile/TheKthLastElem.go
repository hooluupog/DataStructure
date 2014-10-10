// 输入一个单向链表，输出该链表中倒数第k个结点。
// 链表的倒数第0个结点为链表的尾指针。

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

func main() {
	var n, k, i int
	var p, q *LinkList
	fmt.Scan(&n, &k)
	L := new(LinkList)
	CreateList(L, n)
	p = L
	// 设定两个指针，距离相差为k，一个指针走到表尾时，另一个指针恰好为倒数k
	// 位置。
	for i, q = 0, L; i < k; i++ {
		q = q.Next
	}
	// 同步后移两个指针直到一个走到表尾
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next
	}
	fmt.Println(p.Data)
}
