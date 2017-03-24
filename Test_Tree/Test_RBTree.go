package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	op "tree/bintree/operate"
	"tree/bintree/rbtree"
)

func gen() []string {
	s := make([]string, 10000)
	for i, _ := range s {
		s[i] = strconv.Itoa(rand.Intn(30000))
	}
	return s
}

func main() {
	var rbt *rbtree.RBTree
	var s string
	var l []string
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscan(r, &s)
		if err == io.EOF {
			break
		}
		l = append(l, s)
	}
	for _, v := range l {
		rbt = rbt.RBInsert(v)
	}
	fmt.Println("PreOrder:")
	op.PreOrder(rbt)
	fmt.Println()
	fmt.Println("InOrder:")
	op.InOrder(rbt)
	fmt.Println()
	fmt.Println("PostOrder:")
	op.PostOrder(rbt)
	fmt.Println()
	fmt.Println("PreOrder_Stack:")
	op.PreOrder_Stack(rbt)
	fmt.Println()
	fmt.Println("InOrderStack:")
	op.InOrder_Stack(rbt)
	fmt.Println()
	fmt.Println("PostOrder_Stack:")
	op.PostOrder_Stack(rbt)
	fmt.Println()
	fmt.Println("LevelOrder:")
	op.LevelOrder(rbt)
	fmt.Println()
	fmt.Print("Depth: ")
	fmt.Println(op.Depth(rbt))
	st := gen()
	for _, v := range st {
		p := rbt.Search(v)
		if p != nil && p != rbtree.NIL {
			fmt.Fprintln(w, p.Value())
		}
	}
	w.Flush()
}
