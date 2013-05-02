#include <string>
typedef int Status;

#ifndef _BINTREE_H_
#define _BINTREE_H_
#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE
typedef struct BinNode{
	struct BinNode *lchild,*rchild;
	std::string data;
}BinNode,*BinTree;

BinTree Create_BinTree();//Creat a linked storage structure of binary tree with preorder
void Visit(BinTree T);//Print the node being visited now
int ToInt(BinTree T);//Convert numerical string into int value
void PreOrder(BinTree T); //Recusively preorder traverse binary tree
void InOrder(BinTree T); //Recusively inorder traverse binary tree
void PostOrder(BinTree T); //Recusively postorder traverse binary tree
int Depth(BinTree T); 
#endif
