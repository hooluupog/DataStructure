// 从尾到头输出链表
// input 
// 5
// 1 2 3 4 5
// output
// 5 4 3 2 1

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

func InvertPrint(L *LinkList) {
	if L == nil {
		return
	}
	InvertPrint(L.Next)
	fmt.Printf("%d ", L.Data)
}

func main() {
	var n int
	fmt.Scan(&n)
	L := new(LinkList)
	CreateList(L, n)
	InvertPrint(L.Next)
}
