// Print the path whose nodes weight's sum equals e in bintree.
package main

import (
	"bintree"
	"fmt"
)

var path string
var sum int

//路径从根结点开始向下一直到叶节点所经过的结点形成的路径
func findPath(T *bintree.BinTree, e int) {
	if T != nil {
		b := []byte(T.Data)
		sum += int(b[0])
		path += T.Data
		if T.Lchild == nil && T.Rchild == nil && sum == e {
			fmt.Println(path)
		}
		findPath(T.Lchild, e)
		findPath(T.Rchild, e)
		sum -= int(b[0])
		path = path[:(len(path) - 1)]
	}
}

// 路径不必从根结点开始,不必在叶节点结束,开始和结束结点一个为另一个的祖先
func findPath2(T *bintree.BinTree, e int, level int) {
	if T != nil {
		sum := 0
		path += T.Data
		b := []byte(path)
		for i := level; i > -1; i-- {
			sum += int(b[i])
			if sum == e {
				fmt.Println(path[i:len(path)])
			}
		}
		findPath2(T.Lchild, e, level+1)
		findPath2(T.Rchild, e, level+1)
		path = path[:(len(path) - 1)]
	}
}

func main() {
	var t *bintree.BinTree
	t = t.Create_BinTree()
	fmt.Println("先序序列：")
	t.PreOrder()
	fmt.Println("\n中序序列：")
	t.InOrder()
	//e := 213
	e := 680
	fmt.Printf("\n结点权值之和等于%d的路径：\n", e)
	findPath(t, e)
	findPath2(t, e, 0)
	//go run FindPathSum.go<sum.txt
}
