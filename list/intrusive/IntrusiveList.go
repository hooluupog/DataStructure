/** Intrusive LinkedList implementation.
*  Definition:
*      An intrusive data structure is one where the data items contain the metadata needed,
*      instead of the metadata pointing to the contained data.The intrusive linked list is
*      the data oriented(FYI,data oriented programming) version of the linked list container.
*      For example,the external linked list looks like this:
*
*      type Elem struct{
*           val interface{}
*           next *Elem
*      }
*      type List struct{
*           first *Elem
*      }
*
*             -> Elem.next -> Elem.next -> nil
*                .obj         .obj
*                  |            |
*                  v            v
*                 data         data
*
*     Whereas, The Intrusive linked list:
*     type Elem struct{
*         pre  *List
*         next *List
*     }
*     type List struct{
*         first *Elem
*     }
*
*             -> data
*                    .Elem.next -> data
*                                      .Elem.next -> nil
*
*     As you see, the 'intrusive' means that the data contains Elem.
*     your data 'T' looks like this,
*     type T struct{
*          val interface{}
*          intrusive.List // embedded a linklist.
*     }
*     The goal of this is to decrease the number of allocations; instead of `Elem`
*     and `data` being separate chunks of memory, they are allocated as one.
*     For more details,see http://www.codefarms.com/publications/intrusiv/intr.htm.
*
*  Pro:
*     Decoupling memory allocation from the container itself to avoid unnecessary
*     memory copies;
*     Circular implementation simplifies list manipulation implementation.
* */

package intrusive

import (
	"fmt"
)

type Elem struct {
	list *List
	pre  item
	next item
}

func (e *Elem) Prev() item {
	if p := e.pre; e.list != nil && p != e.list.first {
		return p
	}
	return nil
}

func (e *Elem) Next() item {
	if p := e.next; e.list != nil && p != e.list.first {
		return p
	}
	return nil
}

type List struct {
	first  item
	length int
}

func (l *List) IsEmpty() bool {
	return l.length == 0
}

func (l *List) Len() int {
	return l.length
}

type item interface {
	Get() *Elem
}

func (l *List) First() item { // return head node.
	if l.Len() == 0 {
		return nil
	}
	return l.first
}

func (l *List) Last() item {
	if l.Len() == 0 {
		return nil
	}
	return l.first.Get().pre
}

func (l *List) Add(elem item) { // append [elem] into LinkList.
	l.insertBefore(l.first, elem, false)
}

func (l *List) AddFirst(elem item) { // add [elem] at the head of LinkList.
	l.insertBefore(l.first, elem, true)
}

func (l *List) insertBefore(next item, elem item, updateFirst bool) {
	if elem.Get().list != nil {
		panic("Element is already in a list.")
	}
	elem.Get().list = l // Link elem into list.
	if l.IsEmpty() {
		elem.Get().pre = elem
		elem.Get().next = elem
		l.first = elem
	} else {
		prev := next.Get().pre
		next.Get().pre = elem
		elem.Get().pre = prev
		elem.Get().next = next
		prev.Get().next = elem
		if updateFirst {
			l.first = elem
		}
	}
	l.length++
}

// Drop [elem] 's link from LinkList.
func (l *List) unlink(elem item) {
	e := elem.Get()
	next := e.next
	e.pre.Get().next = e.next
	e.next.Get().pre = e.pre
	e.pre = e.next
	e.list = nil
	l.length--
	if l.IsEmpty() {
		l.first = nil
	} else if elem == l.first {
		l.first = next
	}
}

func (l *List) Remove(elem item) bool {
	e := elem.Get()
	if e.list != l {
		return false
	}
	e.list = nil // unlink elem from list.
	l.unlink(elem)
	return true
}

func (l *List) ReverseBetween(m, n int) {
	if m < 1 || n > l.length {
		panic(fmt.Sprintf("Index out of bounds.Wrong range (%d,%d).", m, n))
	}
	cnt := n - m
	recovery := false
	head := l.first
	p := l.first
	for i := 1; i < m; i++ { // p points to the first node to be reversed.
		p = p.Get().next
	}
	if p != l.first {
		recovery = true // Need do head node recovery
	}
	next := p
	for cnt > 0 { // Insert nodes from m to n into list.
		q := p.Get().next
		l.unlink(q)
		l.insertBefore(next, q, true)
		next = q
		cnt--
	}
	if recovery {
		l.first = head // head node recovery.
	}
}
