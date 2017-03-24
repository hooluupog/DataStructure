package main

import (
	"fmt"
	"sort"
)

type Color bool

const (
	Red   = true
	Black = false
)

var NIL = &RBTree{color: Black} // Sentinel nodes.

type RBTree struct {
	val                 byte
	idx                 []int // Duplicate elements index stored as sorted array.
	left, right, parent *RBTree
	color               Color
}

func levelOrder(root *RBTree) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var data []*RBTree
	var queue []*RBTree
	level := make(map[*RBTree]int)
	queue = append(queue, root)
	level[root] = 0
	rows := 0
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		data = append(data, q)
		if rows < level[q] {
			rows = level[q]
		}
		if q.left != NIL {
			queue = append(queue, q.left)
			level[q.left] = level[q] + 1
		}
		if q.right != NIL {
			queue = append(queue, q.right)
			level[q.right] = level[q] + 1
		}
	}
	res = make([][]int, rows+1)
	for _, v := range data {
		res[level[v]] = append(res[level[v]], int(v.val))
	}
	return res
}
func InSort(value int, b []int) []int { //Insertion sort.
	pos := sort.Search(len(b), func(i int) bool { return b[i] >= value })
	if pos == len(b) {
		b = append(b, value)
	} else {
		b = append(b, b[len(b)-1])
		for i := len(b) - 2; i > pos; i-- {
			b[i] = b[i-1]
		}
		b[pos] = value
	}
	return b
}
func (t *RBTree) RBInsert(b byte, i int) *RBTree { // Generate Red-Black tree.
	p, q := t, NIL
	for p != nil && p != NIL { // Binary Search Tree insert.
		q = p
		if b < p.val {
			p = p.left
		} else if b > p.val {
			p = p.right
		} else {
			p.idx = InSort(i, p.idx)
			break
		}
	}
	rb := &RBTree{val: b, color: Red, parent: q}
	if q == NIL {
		t = rb
	} else {
		if b < q.val {
			q.left = rb
		}
		if b > q.val {
			q.right = rb
		}
	}
	rb.left = NIL
	rb.right = NIL
	rb.idx = InSort(i, rb.idx)
	t = t.RBInsertFixup(rb)
	return t
}
func (t *RBTree) LeftRotate(p *RBTree) *RBTree {
	q := p.right
	p.right = q.left
	q.left.parent = p
	q.parent = p.parent
	if p.parent == NIL { // root node.
		t = q
	} else if p == p.parent.left {
		p.parent.left = q
	} else {
		p.parent.right = q
	}
	q.left = p
	p.parent = q
	return t
}
func (t *RBTree) RightRotate(p *RBTree) *RBTree {
	q := p.left
	p.left = q.right
	q.right.parent = p
	q.parent = p.parent
	if p.parent == NIL { // root node.
		t = q
	} else if p == p.parent.left {
		p.parent.left = q
	} else {
		p.parent.right = q
	}
	q.right = p
	p.parent = q
	return t
}
func (t *RBTree) RBInsertFixup(p *RBTree) *RBTree { //  Adjust binary search tree into red-black tree.
	for p.parent.color == Red {
		if p.parent == p.parent.parent.left {
			uncle := p.parent.parent.right
			if uncle.color == Red { // case 1.
				p.parent.color = Black
				uncle.color = Black
				p.parent.parent.color = Red
				p = p.parent.parent
			} else {
				if p == p.parent.right { // case 2.
					p = p.parent
					t = t.LeftRotate(p)
				}
				// case 3.
				p.parent.color = Black
				p.parent.parent.color = Red
				t = t.RightRotate(p.parent.parent)
			}
		} else { //same as above  with "right" and "left" exchanged.
			uncle := p.parent.parent.left
			if uncle.color == Red { // case 1.
				p.parent.color = Black
				uncle.color = Black
				p.parent.parent.color = Red
				p = p.parent.parent
			} else {
				if p == p.parent.left { // case 2.
					p = p.parent
					t = t.RightRotate(p)
				}
				// case 3.
				p.parent.color = Black
				p.parent.parent.color = Red
				t = t.LeftRotate(p.parent.parent)
			}
		}
	}
	t.color = Black
	return t
}
func (t *RBTree) Search(b byte, i int) int { //Return b's index which is bigger than i.
	//If b is not found,return -1.
	if t == nil {
		return -1
	}
	for t != NIL {
		if b < t.val {
			t = t.left
		} else if b > t.val {
			t = t.right
		} else {
			pos := sort.Search(len(t.idx), func(k int) bool { return t.idx[k] > i })
			if pos < len(t.idx) {
				return t.idx[pos]
			}
			break
		}
	}
	return -1
}
func isSubsequence_usingHash(s string, t string) bool {
	var dict [26][]int
	for i, v := range t {
		m := v - 97
		dict[m] = InSort(i, dict[m])
	}
	idx := -1
	for _, v := range s {
		m := v - 97
		pos := sort.Search(len(dict[m]), func(i int) bool { return dict[m][i] > idx })
		if pos == len(dict[m]) {
			return false
		} else {
			idx = dict[m][pos]
		}
	}
	return true
}
func isSubsequence_usingRBTree(s string, t string) bool {
	var T *RBTree
	for i, v := range t {
		T = T.RBInsert(byte(v), i)
	}
	pre := -1 //last founded node index.
	for _, v := range s {
		b := byte(v)
		idx := T.Search(b, pre)
		if idx <= pre {
			return false
		} else {
			pre = idx
		}
	}
	return true
}

func main() {
	s := "leeeeetcode"
	t := "yyylyyyyeyyyyeyyyyyyeyyyyytyyyyycyyyyyyyoyyyyyydyyyyyyyeyyy"
	fmt.Println(isSubsequence_usingHash(s, t))
	fmt.Println(isSubsequence_usingRBTree(s, t))
}
