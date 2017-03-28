type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	cnt := n - m + 1
	var L, pre, p, q *ListNode // L is head node of list which is empty node.
	L = &ListNode{}
	pre = L
	p = head
	for i := 1; i < m; i++ { // p point to the first node to be reversed.
		pre = p
		p = p.Next
	}
	reversedEndNode := p // The last node of reversed subList.
	pre.Next = nil       // pre is head node of subList to be reversed.
	for cnt > 0 {        // insert nodes to head of list.
		temp := pre.Next
		pre.Next = p
		q = p.Next
		p.Next = temp
		p = q
		cnt--
	}
	if p != nil { // Connect reversed subList to origin list.
		reversedEndNode.Next = p
	}
	if L != pre { // Find the head node.
		L.Next = head
	}
	return L.Next
}
