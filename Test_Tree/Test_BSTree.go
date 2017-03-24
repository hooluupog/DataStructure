package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"tree/bintree/bstree"
	op "tree/bintree/operate"
)

func gen() []string {
	s := make([]string, 10000)
	for i, _ := range s {
		s[i] = strconv.Itoa(rand.Intn(30000))
	}
	return s
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var bst *bstree.BSTree
	var s string
	var l []string
	for {
		_, err := fmt.Fscan(r, &s)
		if err == io.EOF {
			break
		}
		l = append(l, s)
	}
	for _, v := range l {
		bst = bst.Insert(v)
	}
	fmt.Println("PreOrder:")
	op.PreOrder(bst)
	fmt.Println()
	fmt.Println("InOrder:")
	op.InOrder(bst)
	fmt.Println()
	fmt.Println("PostOrder:")
	op.PostOrder(bst)
	fmt.Println()
	fmt.Println("PreOrder_Stack:")
	op.PreOrder_Stack(bst)
	fmt.Println()
	fmt.Println("InOrderStack:")
	op.InOrder_Stack(bst)
	fmt.Println()
	fmt.Println("PostOrder_Stack:")
	op.PostOrder_Stack(bst)
	fmt.Println()
	fmt.Println("LevelOrder:")
	op.LevelOrder(bst)
	fmt.Println()
	fmt.Print("Depth: ")
	fmt.Println(op.Depth(bst))
	st := gen()
	for _, v := range st {
		p := bst.Search(v)
		if p != nil {
			fmt.Fprintln(w, p.Value())
		}
	}
	w.Flush()
}
