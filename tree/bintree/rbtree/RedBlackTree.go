/**
* Title: Red-Black Tree
* Time: 2017/3/24
* Author: Lee
* Details:
* 红黑树的5条性质：
* 1）每个结点要么是红色，要么是黑色;
* 2）根结点为黑色;
* 3）每个叶结点（叶结点为树尾端的空结点，用NIL指针标记）为黑色;
* 4）父结点是红色，则它的孩子结点都是黑色;
* 5）对任一结点，其到叶结点(树尾端NIL指针)的每一条路径都包含相同数目的黑色结点。
**/

package rbtree

import (
	op "tree/bintree/operate"
)

type Color bool

const (
	Red   = true
	Black = false
)

var NIL = &RBTree{color: Black} // Sentinel nodes.

type RBTree struct {
	val                 string
	left, right, parent *RBTree
	color               Color
}

func (t *RBTree) RBInsert(s string) *RBTree { // Generate Red-Black tree.
	val := op.ToInt(s)
	p, q := t, NIL
	for p != nil && p != NIL { // Binary Search Tree insert.
		q = p
		pval := op.ToInt(p.val)
		if val < pval {
			p = p.left
		} else if val > pval {
			p = p.right
		} else {
			break
		}
	}
	rb := &RBTree{val: s, color: Red, parent: q}
	if q == NIL {
		t = rb
	} else {
		qval := op.ToInt(q.val)
		if val < qval {
			q.left = rb
		}
		if val > qval {
			q.right = rb
		}
	}
	rb.left = NIL
	rb.right = NIL
	t = t.RBInsertFixup(rb) //修正使其重新满足红黑树的性质
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
func (t *RBTree) Search(s string) *RBTree {
	val := op.ToInt(s)
	for t != nil && t != NIL {
		tval := op.ToInt(t.val)
		if val < tval {
			t = t.left
		} else if val > tval {
			t = t.right
		} else {
			return t
		}
	}
	return nil
}

func (t *RBTree) Value() string {
	return t.val
}

func (t *RBTree) Left() op.Interface {
	return t.left
}

func (t *RBTree) Right() op.Interface {
	return t.right
}

func (t *RBTree) IsNil() bool {
	return t == nil
}
