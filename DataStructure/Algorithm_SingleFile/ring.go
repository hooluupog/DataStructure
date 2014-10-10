// 检测一个较大的单链表是否有环

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
	//s.Next = L.Next.Next.Next // Generate the ring.
}

func HasRing(L *LinkList, n int) bool {
	mp := make(map[*LinkList]int, n)
	for p := L.Next; p != nil; p = p.Next {
		if mp[p] == 0 { //第一次访问该结点，mp[p] = 1
			mp[p] = 1
		} else { // 有环
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	L := new(LinkList)
	CreateList(L, n)
	fmt.Println(HasRing(L, n))
}
