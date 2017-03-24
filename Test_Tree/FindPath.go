// Find path from m to n in bintree.
package main

import (
	"bintree"
	"fmt"
)

var path string
var record int

func findPath(T *bintree.BinTree, m string, n string) {
	if T != nil {
		if T.Data == m { //start record path
			record = 1
		}
		if T.Data == n { // stop record path
			path += n
			record = 0
		}
		if record == 1 {
			path += T.Data
		}
		findPath(T.Lchild, m, n)
		findPath(T.Rchild, m, n)
		if record == 1 {
			path = path[:(len(path) - 1)]
		}
	}
}

func main() {
	var t *bintree.BinTree
	t = t.Create_BinTree()
	fmt.Println("先序序列：")
	t.PreOrder()
	fmt.Println("\nM到N的路径：")
	findPath(t, "M", "N")
	fmt.Println(path)
}
