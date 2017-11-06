package main

import (
	"fmt"
	"strings"
    "tree/bintree/bintree"
    op "tree/bintree/operate"
)

type BinTree = bintree.BinTree

func serilization(t *BinTree) string {
	if t == nil {
		return "#"
	}
	return fmt.Sprintf("%s", t.Data) + "," + serilization(t.Lchild) + "," + serilization(t.Rchild)
}

func deserilization(seq *[]string) *BinTree {
	val := (*seq)[0]
	*seq = (*seq)[1:]
	if val == "#" {
		return nil
	}
	root := &BinTree{Data: val}
	root.Lchild = deserilization(seq)
	root.Rchild = deserilization(seq)
	return root
}

func main() {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	//str2 := "1,2,4,#,#,#,3,2,4,#,#,#,4,#,#";
	s := strings.Split(str, ",")
	T := deserilization(&s)
	op.PreOrder(T)
    fmt.Println()
    fmt.Println(serilization(T))
}
