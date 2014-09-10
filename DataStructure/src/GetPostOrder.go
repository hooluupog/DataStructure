/*
* input preOrder and inOrder list of a binary tree,
* output its postOrder list.
* 思路：
* 1）先序序列第一个元素为根结点，将线性表分为两部分（即二叉树的左右子树）
* ，再分别处理两个部分找出相应的先序和中序序列。
* 2）对每个部分再执行1），直到剩余一个结点，按照LRN的顺序输出。
* 上面是个递归的过程，最终会得到一个后序序列.
 */

package main

import "fmt"

var illegal bool

func PostOrder(preOrder, inOrder string) string {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return ""
	}
	root := preOrder[0]
	i := 0
	for i < len(inOrder) && inOrder[i] != root {
		i++
	}
	if i == len(inOrder) {
		illegal = true
		return ""
	}
	return PostOrder(preOrder[1:i+1], inOrder[0:i]) +
		PostOrder(preOrder[i+1:len(preOrder)], inOrder[i+1:len(inOrder)]) + string(root)

}
func main() {
	var preOrder, inOrder string
	fmt.Scan(&preOrder, &inOrder)
	illegal = false
	res := PostOrder(preOrder, inOrder)
	if illegal {
		fmt.Println("illegal input")
	} else {
		fmt.Println(res)
	}
}
