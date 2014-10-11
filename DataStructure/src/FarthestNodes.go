// 二叉树中最远的两个结点
// 树中最远结点分为两种情况：1.不经过根结点；2，经过根结点
// 设第 K 棵子树中相距最远的两个节点：Uk和Vk，
// 其距离定义为d（Uk, Vk），那么节点 Uk或Vk即为子树K
// 到根节点Rk 距离最长的节点。不失一般性，我们设Uk为子树K中到根
// 节点Rk 距离最长的节点，其到根节点的距离定义为d（Uk, R）。
// 取d（Ui, R）（1≤i≤k）中最大的两个值max1 和max2，
// 那么经过根节点R的最长路径为max1+max2+2，所以树R中相距最远的两
// 个点的距离为：max{d（U1, V1）, …, d（Uk, Vk），max1+max2+2}。

package main

import (
	"bufio"
	"fmt"
	"os"
)

type BinTree struct {
	Lchild    *BinTree
	Rchild    *BinTree
	MaxLeft   int // 左子树中的最长距离
	MaxRight  int // 右子树中的最长距离
	Data      string
	LeftNode  string // 左子树中的最远结点
	RightNode string // 右子树中的最远结点
}

type Distance struct {
	Nodes    string
	Distance int
}

var r = bufio.NewReader(os.Stdin)

func CreateBinTree(t *BinTree) (v *BinTree) {
	var e string
	fmt.Fscanf(r, "%s", &e)
	r.ReadString('\n') // clean stdin buffer
	if e == "#" {
		t = nil
	} else {
		t = new(BinTree)
		t.Data = e
		t.Lchild = CreateBinTree(t.Lchild)
		t.Rchild = CreateBinTree(t.Rchild)
	}
	v = t
	return v
}

func FindMaxDistance(t *BinTree, d *Distance) {
	if t == nil { // 遍历到叶子结点，返回
		return
	}
	if t.Lchild == nil { // 如果左子树为空，该结点的左边最长距离为0
		t.MaxLeft = 0
		t.LeftNode = t.Data
	}
	if t.Rchild == nil { // 如果右子树为空，该结点的右边最长距离为0
		t.MaxRight = 0
		t.RightNode = t.Data
	}
	if t.Lchild != nil { // 左子树不空，递归寻找左子树最长距离
		FindMaxDistance(t.Lchild, d)
	}
	if t.Rchild != nil { // 右子树 不空，递归寻找右子树最长距离
		FindMaxDistance(t.Rchild, d)
	}
	// 计算左子树最长结点距离
	if t.Lchild != nil {
		tempMax := 0
		if t.Lchild.MaxLeft > t.Lchild.MaxRight {
			tempMax = t.Lchild.MaxLeft
			t.LeftNode = t.Lchild.LeftNode
		} else {
			tempMax = t.Lchild.MaxRight
			t.LeftNode = t.Lchild.RightNode
		}
		t.MaxLeft = tempMax + 1
	}
	// 计算右子树最长结点距离
	if t.Rchild != nil {
		tempMax := 0
		if t.Rchild.MaxLeft > t.Rchild.MaxRight {
			tempMax = t.Rchild.MaxLeft
			t.RightNode = t.Rchild.LeftNode
		} else {
			tempMax = t.Rchild.MaxRight
			t.RightNode = t.Rchild.RightNode
		}
		t.MaxRight = tempMax + 1
	}
	// 更新最长距离
	if t.MaxLeft+t.MaxRight > d.Distance {
		d.Distance = t.MaxLeft + t.MaxRight
		d.Nodes = t.LeftNode + " " + t.RightNode
	}
}

func main() {
	var t *BinTree
	t = CreateBinTree(t)
	d := &Distance{}
	FindMaxDistance(t, d)
	fmt.Println("最远结点:", d.Nodes, "\n距离:", d.Distance)
}
