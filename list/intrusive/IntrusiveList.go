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
*           head *Elem
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
*         head *Elem
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

type item interface {
	Get() *Elem
}

type Elem struct {
	list *List
	pre  item
	next item
}

func (e *Elem) Get() *Elem { return e }

func (e *Elem) Prev() item {
	if p := e.pre; e.list != nil && p != e.list.head {
		return p
	}
	return nil
}

func (e *Elem) Next() item {
	if p := e.next; e.list != nil && p != e.list.head {
		return p
	}
	return nil
}

type List struct {
	head item
	len  int
}

func (l *List) Init() *List { // Init list l with head node.
	elem := Elem{list: l}
	l.head = &elem
	elem.pre, elem.next = l.head, l.head
	l.len = 0
	return l
}

func New() *List { return new(List).Init() }

func (l *List) Len() int {
	return l.len
}

func (l *List) First() item { // return first node.
	if l.Len() == 0 {
		return nil
	}
	return l.head.Get().next
}

func (l *List) Last() item {
	if l.Len() == 0 {
		return nil
	}
	return l.head.Get().pre
}

func (l *List) Add(elem item) { // append [elem] into LinkList.
	l.insert(elem, l.head.Get().pre)
}

func (l *List) AddFirst(elem item) { // add [elem] at the first.
	l.insert(elem, l.head)
}

func (l *List) InsertBefore(elem item, at item) { //Insert [elem] before [at].
	l.insert(elem, at.Get().pre)
}

func (l *List) InsertAfter(elem item, at item) { //Insert [elem] after [at].
	l.insert(elem, at)
}

func (l *List) insert(elem item, at item) {
	if elem.Get().list != nil {
		panic("Element is already in a list.")
	}
	e, a := elem.Get(), at.Get()
	e.list = l // Link elem into list.
	e.pre = at
	e.next = a.next
	a.next = elem
	e.next.Get().pre = elem
	l.len++
}

func (l *List) Remove(elem item) item {
	e := elem.Get()
	e.list = nil // unlink elem from list.
	e.pre.Get().next = e.next
	e.next.Get().pre = e.pre
	e.pre = nil
	e.next = nil
	e.list = nil
	l.len--
	return elem
}

func (l *List) ReverseBetween(m, n int) {
	if m < 1 || n > l.Len() {
		panic(fmt.Sprintf("Index out of bounds.Wrong range (%d,%d).", m, n))
	}
	cnt := n - m
	start := l.First()
	for i := 1; i < m; i++ { //start points to the first node to be reversed.
		start = start.Get().Next()
	}
	end := start
	for cnt > 0 { // Insert nodes from m to n into list.
		curr := end.Get().Next()
		l.Remove(curr)
		l.InsertBefore(curr, start)
		start = curr
		cnt--
	}
}
