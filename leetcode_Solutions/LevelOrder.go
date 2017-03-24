package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

var Stdin = bufio.NewReader(os.Stdin)

func (t *TreeNode) CreateTree() (v *TreeNode) {
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
		t = new(TreeNode)
		t.Val, _ = strconv.Atoi(e)
		t.Left = t.CreateTree()
		t.Right = t.CreateTree()
	}
	v = t
	return
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var data []*TreeNode
	var queue []*TreeNode
	level := make(map[*TreeNode]int)
	queue = append(queue, root)
	level[root] = 0
	rows := 0
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		data = append(data, q)
		if rows < level[q] {
			rows = level[q]
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
			level[q.Left] = level[q] + 1
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
			level[q.Right] = level[q] + 1
		}
	}
	res = make([][]int, rows+1)
	for _, v := range data {
		res[level[v]] = append(res[level[v]], v.Val)
	}
	return res
}
func main() {
	var T *TreeNode
	T = T.CreateTree()
	fmt.Println(levelOrder(T))
}
