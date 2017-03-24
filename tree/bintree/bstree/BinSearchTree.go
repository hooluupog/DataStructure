package bstree

import (
	op "tree/bintree/operate"
)

type BSTree struct {
	val         string
	left, right *BSTree
}

func (t *BSTree) Insert(s string) *BSTree {
	var p, q *BSTree
	val := op.ToInt(s)
	p, q = t, nil
	for p != nil {
		q = p
		pval := op.ToInt(p.val)
		if val < pval {
			p = p.left
		} else if val > pval {
			p = p.right
		} else {
			return t
		}
	}
	if q == nil {
		t = &BSTree{val: s}
		return t
	}
	qval := op.ToInt(q.val)
	if val < qval {
		q.left = &BSTree{val: s}
	}
	if val > qval {
		q.right = &BSTree{val: s}
	}
	return t
}

func (t *BSTree) Search(s string) *BSTree {
	val := op.ToInt(s)
	for t != nil {
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

func (t *BSTree) Value() string {
	return t.val
}

func (t *BSTree) Left() op.Interface {
	return t.left
}

func (t *BSTree) Right() op.Interface {
	return t.right
}

func (t *BSTree) IsNil() bool {
	return t == nil
}
