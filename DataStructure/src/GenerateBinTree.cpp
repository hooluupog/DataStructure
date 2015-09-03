/*
 * 已知中序、后序序列，生成2叉树；
 * 已知先序、中序序列，生成2叉树
 *
 */
#include "stdio.h"
#include "stdlib.h"
typedef struct BiNode {
	int data;
	struct BiNode *lchild, *rchild;
} BiNode, *BiTree;

//Print the node being visited now
void Visit(BiTree T) {
	printf("%d ", T->data);
}

//Recusively preorder traverse binary tree
void PreOrder(BiTree T) {
	if (T != NULL) {
		Visit(T);
		PreOrder(T->lchild);
		PreOrder(T->rchild);
	}
}

BiNode *BuildFromPreInOrder(int *pre, int *in, int ps, int pe, int is, int ie) {
	// 已知先序、中序序列，生成2叉树
	if (ps > pe || is > ie) // 子序列长度小于1
		return NULL;
	BiTree r;
	int i = is;
	while (pre[ps] != in[i++]); // 先序序列第一个元素为根结点，在中序序列中寻找根结点的位置
	int llen = i - is - 1; // 左子序列长度
	int rlen = ie - i + 1; // 右子序列长度
	r = (BiTree) malloc(sizeof(BiNode));
	r->data = pre[ps];
	r->lchild = BuildFromPreInOrder(pre, in, ps + 1, ps + llen, is,is + llen - 1);
	r->rchild = BuildFromPreInOrder(pre, in, pe - rlen + 1, pe, ie - rlen + 1,ie);
	return r;
}

BiNode *BuildFromPostInOrder(int *post, int *in, int ps, int pe, int is,int ie) {
	// 已知中序、后序序列，生成2叉树
	if (ps > pe || is > ie)
		return NULL;
	BiTree r;
	int i = is;
	while (post[pe] != in[i++]);// 中序序列最后一个元素为根结点，在中序序列中寻找根结点的位置
	int llen = i - is - 1; // 左子序列长度
	int rlen = ie - i + 1; // 右子序列长度
	r = (BiTree) malloc(sizeof(BiNode));
	r->data = post[pe];
	r->lchild = BuildFromPostInOrder(post, in, pe - llen, pe - 1, is,is + llen - 1);
	r->rchild = BuildFromPostInOrder(post, in, ps, ps + rlen - 1, ie - rlen + 1, ie);
	return r;
}

void testGenerate() {
	//       1
    //     /   \
	//    2     3
	//   / \   /
	//  4   5 6
	int pre[6] = { 1, 2, 4, 5, 3, 6 };
	int in[6] = { 4, 2, 5, 1, 6, 3 };
	int post[6] = { 6, 3, 5, 4, 2, 1 };
	BiTree r;
	r = BuildFromPreInOrder(pre, in, 0, 5, 0, 5);
	PreOrder(r);
	printf("\n");
	r = BuildFromPostInOrder(post, in, 0, 5, 0, 5);
	PreOrder(r);
	printf("\n");
}
