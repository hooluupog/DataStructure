// Invert binary tree
//    4                4
//   / \              / \
//  2   7      to    7   2
//     / \          / \
//    1   3        3   1

package main

import (
	"bintree"
	"fmt"
	"stack"
)

func invert(root *bintree.BinTree) {
	if root == nil {
		return
	}
	s := new(stack.Stack)
	s.Push(root)
	for s.Length() != 0 {
		t := s.Pop().(*bintree.BinTree)
		t.Lchild, t.Rchild = t.Rchild, t.Lchild
		if t.Lchild != nil {
			s.Push(t.Lchild)
		}
		if t.Rchild != nil {
			s.Push(t.Rchild)
		}
	}
}
func main() {
	var t *bintree.BinTree
	t = t.Create_BinTree()
	t.PreOrder()
	fmt.Println()
	invert(t)
	t.PreOrder()
}
