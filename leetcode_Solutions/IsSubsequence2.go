package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type BST struct {
	val         byte
	idx         []int // Duplicate elements index.
	left, right *BST
}

func levelOrder(root *BST) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var data []*BST
	var queue []*BST
	level := make(map[*BST]int)
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
		if q.left != nil {
			queue = append(queue, q.left)
			level[q.left] = level[q] + 1
		}
		if q.right != nil {
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
func InSort(value int, b []int) []int {
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
func (t *BST) BinSearchTree(s string) *BST { // todo: Balanced BST.
	if len(s) == 0 {
		return nil
	}
	t = &BST{val: byte(s[0])}
	t.idx = append(t.idx, 0)
	for i, v := range s {
		b := byte(v)
		p, q := t, t
		for p != nil {
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
		if b < q.val {
			q.left = &BST{val: b}
			q.left.idx = InSort(i, q.left.idx)
		}
		if b > q.val {
			q.right = &BST{val: b}
			q.right.idx = InSort(i, q.right.idx)
		}
	}
	return t
}
func (t *BST) Search(b byte, i int) int { //Return b's index which is bigger than i.
	//If b is not found,return -1.
	for t != nil {
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
func isSubsequence_usingBST(s string, t string) bool {
	var T *BST
	T = T.BinSearchTree(t)
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
func gen() []byte {
	N := 1 + rand.Intn(26)
	b := make([]byte, N)
	for i := 0; i < N; i++ {
		b[i] = byte(97 + rand.Intn(26))
	}
	return b
}

func main() {
	s := "leeeeetcode"
	t := "yyylyyyyeyyyyeyyyyyyeyyyyytyyyyycyyyyyyyoyyyyyydyyyyyyyeyyy"
	fmt.Println(isSubsequence_usingHash(s, t))
	fmt.Println(isSubsequence_usingBST(s, t))
}
