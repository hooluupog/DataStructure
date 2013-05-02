/*链式结构存储的二叉树*/

#include "BinTree.h"
#include "SqStack.h"
#include "fatal.h"
#include "stdio.h"
#include <iostream>
#include <string>
#include "string.h"
using namespace std;

//Creat a linked storage structure of binary tree with preorder
BinTree Create_BinTree() {
	//The legal inputs are numerical string and "#"("#" represents empty node)
	BinTree L;
	string e;
	cin >> e;
	if(e == "#") {
		L = NULL;
	}
	else {
		L = new BinNode;
		L->data = e;
		L->lchild = Create_BinTree();
		L->rchild = Create_BinTree();
	}
	return L;
}

//Print the node being visited now
void Visit(BinTree T) {
	cout << T->data << " ";
}

//Convert numerical string into int value
int ToInt(BinTree T){
	char buf[12];
	int value;
	strcpy(buf,T->data.c_str());
	sscanf(buf,"%d",&value);
	return value;
}

//Recusively preorder traverse binary tree
void PreOrder(BinTree T) {
	if(T != NULL){
		Visit(T);
		PreOrder(T->lchild);
		PreOrder(T->rchild);
	}
}

//Recusively inorder traverse binary tree
void InOrder(BinTree T) {
	if(T != NULL){
		InOrder(T->lchild);
		Visit(T);
		InOrder(T->rchild);
	}
}

//Recusively postorder traverse binary tree
void PostOrder(BinTree T) {
	if(T != NULL){
		PostOrder(T->lchild);
		PostOrder(T->rchild);
		Visit(T);
	}
}

int Depth(BinTree T) {
	if(T == NULL) return 0;
	else {
		int ldepth,rdepth;
		return (ldepth = Depth(T->lchild)) > (rdepth = Depth(T->rchild)) ? (ldepth + 1) : (rdepth + 1);
	}
}
