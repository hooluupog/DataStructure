#include "BinTree.h"
#include "stdio.h"
#include <iostream>
using namespace std;

void testBinTree() {
	freopen("input/BinTree.txt","r",stdin);
	printf("BinTree test: \n");
	BinTree T;
	T = Create_BinTree();
	cout << "Preorder:" << endl;
	PreOrder(T);
	cout << "\nInorder:" << endl;
	InOrder(T);
	cout << "\nPostorder:" << endl;
	PostOrder(T);
	cout << "\nDepth: " << Depth(T) << endl;
}
