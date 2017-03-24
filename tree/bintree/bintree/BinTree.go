/**
* Title: Binary Tree abstract Datasture
* Time: 2017/3/24
* Author: Lee
**/

package bintree

import (
	"bufio"
	"fmt"
	"os"
	op "tree/bintree/operate"
)

type BinTree struct {
	Lchild, Rchild *BinTree
	Data           string
}

var Stdin = bufio.NewReader(os.Stdin)

func (t *BinTree) CreateBinTree() (v *BinTree) {
	/*
	** Recusively create a binary tree.The legal inputs are numerical string or "#"("#" represents empty node). The input sequence is as follows:
	** A binary tree is :
	** ----1----
	** --2---3--
	** the input is : 1 2 # # 3 # #
	 */
	var e string
	fmt.Fscanf(Stdin, "%s", &e)
	Stdin.ReadString('\n') // clean stdin buffer
	if e == "#" {
		t = nil
	} else {
		t = new(BinTree)
		t.Data = e
		t.Lchild = t.CreateBinTree()
		t.Rchild = t.CreateBinTree()
	}
	v = t
	return
}

func (t *BinTree) Value() string {
	return t.Data
}

func (t *BinTree) Left() op.Interface {
	return t.Lchild
}

func (t *BinTree) Right() op.Interface {
	return t.Rchild
}

func (t *BinTree) IsNil() bool {
	return t == nil
}
