/*
 * 最大路径和；
 * 2叉树最小/大深度；
 * 判断是否对称树(一颗二叉树是否是自己的镜像)；
 *  例如，下图为对称的:
 *
 *     1
 *    / \
 *   2   2
 *  / \ / \
 * 3  4 4  3
 *
 * 非对称的：
 *     1
 *    / \
 *   2   2
 *    \   \
 *    3    3
 *
 * 判断是否平衡2叉树
 */

#include "stdio.h"
#include "BinTree.h"
#include "stdlib.h"

int max(int a, int b) {
	return a > b ? a : b;
}

int min(int a, int b) {
	return a < b ? a : b;
}

int MaxPathSum(BinTree r, int &res) {
	if (r == NULL)
		return 0;
	int left = MaxPathSum(r->lchild, res);  //左子树叶节点到根结点最大路径和
	int right = MaxPathSum(r->rchild, res);  //右子树叶节点到根结点最大路径和
	int sum = max(ToInt(r), ToInt(r) + max(left, right)); // sum记录叶节点到根结点的最大路径和
	res = max(res, ToInt(r) + left + right);
	res = max(res, sum);  //最大路径和
	return sum;
}

int MinDepth(BinTree r) {
	if (r == NULL)
		return 0;
	int left = MinDepth(r->lchild);
	int right = MinDepth(r->rchild);
	if (left == 0)
		return right + 1;
	if (right == 0)
		return left + 1;
	return min(left, right) + 1;
}

int MaxDepth(BinTree r) {
	if (r == NULL)
		return 0;
	int left = MaxDepth(r->lchild);
	int right = MaxDepth(r->rchild);
	return max(left, right) + 1;
}

int checkTree(BinTree r) {
	if (r == NULL)
		return 0;
	int left = checkTree(r->lchild);
	int right = checkTree(r->rchild);
	if (left < 0 || right < 0 || abs(left - right) > 1)
		return -1;
	return max(left, right) + 1;
}
bool IsAvlTree(BinTree r){
	return checkTree(r) >= 0 ? true : false;
}

bool Symmetric(BinNode *left,BinNode *right){
	if(left == NULL && right == NULL) return true;
	else if(left != NULL && right != NULL){
		return (left->data == right->data && Symmetric(left->lchild,right->rchild) && Symmetric(left->rchild,right->lchild));
	}else{
		return false;
	}
}
bool IsSymmetricTree(BinTree r){
 	if (r == NULL) return false;
 	return Symmetric(r->lchild,r->rchild);
}

void testProblems() {
	freopen("input/BinTree2.txt", "r", stdin);
	BinTree r;
	int res = 0;
	r = Create_BinTree();
	MaxPathSum(r, res);
	printf("%d\n", res);
	printf("%d\n", MinDepth(r));
	printf("%d\n", MaxDepth(r));
	printf("%d\n", IsAvlTree(r));
	printf("%d\n", IsSymmetricTree(r));
}
