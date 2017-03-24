/**
* Title: Test Binary Tree
* Time: 2013/4/16
* Author: Lee
**/

package main

import (
	"fmt"
	"tree/bintree/bintree"
	op "tree/bintree/operate"
)

func main() {
	var t *bintree.BinTree
	t = t.CreateBinTree()
	fmt.Println("PreOrder:")
	op.PreOrder(t)
	fmt.Println("\nInOrder:")
	op.InOrder(t)
	fmt.Println("\nPostOrder:")
	op.PostOrder(t)
	fmt.Println("\nPreOrder_Stack:")
	op.PreOrder_Stack(t)
	fmt.Println("\nInOrder_Stack:")
	op.InOrder_Stack(t)
	fmt.Println("\nPostOrder_Stack:")
	op.PostOrder_Stack(t)
	fmt.Println("\nLevelOrder:")
	op.LevelOrder(t)
	fmt.Println("\nDepth:")
	fmt.Println(op.Depth(t))
}
