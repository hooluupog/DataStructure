/**
* Title: Binary Tree abstract Datasture
* content: General operating functions for binary Tree
* Time: 2017/3/24
* Author: Lee
**/

package operate

import (
	"fmt"
	"os"
	"stack"
	"strconv"
)

type Interface interface {
	Value() string
	Left() Interface
	Right() Interface
	IsNil() bool
}

func Visit(t Interface) {
	fmt.Printf("%v ", t.Value())
}

func Depth(t Interface) (depth int) {
	if t.IsNil() {
		return 0
	} else {
		ldepth := Depth(t.Left())
		rdepth := Depth(t.Right())
		if ldepth > rdepth {
			depth = ldepth + 1
		} else {
			depth = rdepth + 1
		}
	}
	return
}

func ToInt(s string) (value int) {
	value, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return
}

func PreOrder(t Interface) {
	if !t.IsNil() {
		Visit(t)
		PreOrder(t.Left())
		PreOrder(t.Right())
	}
}

func InOrder(t Interface) {
	if !t.IsNil() {
		InOrder(t.Left())
		Visit(t)
		InOrder(t.Right())
	}
}

func PostOrder(t Interface) {
	if !t.IsNil() {
		PostOrder(t.Left())
		PostOrder(t.Right())
		Visit(t)
	}
}

func PreOrder_Stack(t Interface) {
	p := t //p point to root node 't'
	s := new(stack.Stack)
	s.Push(p)             //push root node
	for s.Length() != 0 { // while stack isn't empty
		p = s.Pop().(Interface) //pop out
		Visit(p)                //visit it
		if !p.Right().IsNil() {
			s.Push(p.Right()) //push Right
		}
		if !p.Left().IsNil() {
			s.Push(p.Left()) //push Left
		}
	} //for
}

func InOrder_Stack(t Interface) {
	p := t //p point to root node 't'
	s := new(stack.Stack)
	s.Push(p)             //push root node
	for s.Length() != 0 { // while stack isn't empty
		for !p.IsNil() { // traverse left subtree until the child is nil
			p = p.Left()
			s.Push(p)
		}
		s.Pop() //pop out the nil pointer
		if s.Length() != 0 {
			p = s.Pop().(Interface) // Pop out
			Visit(p)                // visit p
			p = p.Right()           // p move right
			s.Push(p)
		} //if
	} //for
}

func PostOrder_Stack(t Interface) {
	// Differentate left return from right return is key point of this algorithm.
	//p node's Tag == 0: p's left subtree returning; 1: p's right subtree returning
	//if returning from left,continue traverse right subtree
	//if returning from right,visit p node
	tag := make(map[Interface]int)
	if !t.IsNil() {
		tag[t] = 0 // set left subtree mark(unvisitted mark)

	}
	p := t //p point to root node 't'
	s := new(stack.Stack)
	for s.Length() != 0 || !p.IsNil() { // while stack isn't empty
		for !p.IsNil() && tag[p] == 0 { //push into stack only if p not nil and tag[p] == 0.
			s.Push(p)
			p = p.Left() //traverse left subtree
			if !p.IsNil() {
				tag[p] = 0 //set unvisitted mark
			}
		} //for
		if s.Length() != 0 {
			p = s.Top().(Interface) // p point to top stack
			if tag[p] == 1 {        // right returning,visit it
				p = s.Pop().(Interface)
				Visit(p)
			} else { // left returning,traverse right subtree
				tag[p] = 1
				p = p.Right()
				if !p.IsNil() {
					tag[p] = 0
				}
			} //else
		} else {
			break
		}
	} //for
}
func LevelOrder(t Interface) {
	if !t.IsNil() {
		var queue []Interface
		queue = append(queue, t)
		for len(queue) != 0 {
			q := queue[0]
			Visit(q)
			queue = queue[1:]
			if !q.Left().IsNil() {
				queue = append(queue, q.Left())
			}
			if !q.Right().IsNil() {
				queue = append(queue, q.Right())
			}
		}
	}
}
