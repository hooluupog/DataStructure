/**/
/* Get postorder sequence from levelorder and inorder sequence.   */
/* Create a binary tree from the levelorder and inorder sequence. */
/* PostOrder the binary tree and then get the postorder sequence. */
/* sample input                                                   */
/* ABCDEFGHIJKLMNOPQRSTUVWXYZ                                     */
/* GDWQXLHBEIRYMACSNTJOZUFKVP                                     */
/* sample output                                                  */
/* G W X Q L H D Y R M I E B S T N Z U O J V P K F C A            */
/**/

package main

import "fmt"

type BinTree struct {
	data   byte
	lchild *BinTree
	rchild *BinTree
}

type QNode struct {
	data interface{}
	next *QNode
}

type Quene struct {
	front *QNode
	rear  *QNode
}

type QElem struct {
	subsequence string   // Subsequence is left or right part of inorder sequence from which root node could be find out.
	lpos        int      // root node index of levelorder sequence
	bnode       *BinTree // The BinNode of root node.
}

func (Q *Quene) Init() {
	q := new(QNode)
	Q.rear = q
	Q.front = Q.rear
}

func (Q *Quene) EnQuene(value interface{}) {
	p := &QNode{data: value}
	Q.rear.next = p
	Q.rear = p
}

func (Q *Quene) DeQuene() (value interface{}) {
	if Q.IsEmpty() {
		value = nil
	} else {
		value = Q.front.next.data
		p := Q.front.next
		Q.front.next = p.next
		if Q.rear == p {
			Q.rear = Q.front
		}
	}
	return value
}

func (Q *Quene) IsEmpty() bool {
	return Q.front == Q.rear
}

func CreateBinTree(tree *BinTree, level string, in string) {
	// Creat a binary tree from leveOrder and inOrder sequence.
	pos := 0 // It moves in the way of level order and record the cuurent position.
	Q := new(Quene)
	Q.Init()
	Q.EnQuene(QElem{in, 0, tree})
	pos++
	for !Q.IsEmpty() {
		node := Q.DeQuene().(QElem)
		subtree := node.subsequence
		npos := node.lpos // Root node position in levelorder sequence.
		t := node.bnode
		t.data = level[npos]
		if pos <= len(level) {
			i := 0
			// Find the position of root node(level[pos]) in inorder sequence.
			for ; i < len(subtree) && level[npos] != subtree[i]; i++ {
			}
			switch {
			case len(subtree) == 1:
				continue
			case i == 0: // none left child
				rt := new(BinTree)
				t.rchild = rt
				Q.EnQuene(QElem{subtree[i+1:], pos, rt})
				pos++
			case i == len(subtree)-1: // none right child
				lt := new(BinTree)
				t.lchild = lt
				Q.EnQuene(QElem{subtree[:i], pos, lt})
				pos++
			case i == len(subtree): // single node tree
				continue
			default: // Both left child and right child exist
				lt := new(BinTree)
				rt := new(BinTree)
				t.rchild = rt
				t.lchild = lt
				Q.EnQuene(QElem{subtree[:i], pos, lt})
				pos++
				Q.EnQuene(QElem{subtree[i+1:], pos, rt})
				pos++
			}
		}
	}
}

func PostOrder(T *BinTree) {
	if T != nil {
		PostOrder(T.lchild)
		PostOrder(T.rchild)
		fmt.Printf("%s ", string(T.data))
	}
}
func main() {
	var level, in string
	t := new(BinTree)
	fmt.Scan(&level, &in)
	fmt.Println(level, in)
	CreateBinTree(t, level, in)
	PostOrder(t)
	fmt.Println()
}
