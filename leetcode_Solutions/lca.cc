#include "stdio.h"
#include "stdlib.h"
#include <iostream>
#include "string.h"
#include <string>
#define NEW(x) (x*)malloc(sizeof(x))

using namespace std;

struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
};

typedef struct TreeNode* Tree;
typedef struct Node {
	Tree val;
	struct Node *next;
}*List, ListNode;

int ToInt(string s) {
	char buf[12];
	int value;
	strcpy(buf, s.c_str());
	sscanf(buf, "%d", &value);
	return value;
}

Tree Create_TreeNode() {
	//The legal inputs are numerical string and "#"("#" represents empty node)
	Tree L;
	string e;
	cin >> e;
	if (e == "#") {
		L = NULL;
	} else {
		L = new TreeNode;
		L->val = ToInt(e);
		L->left = Create_TreeNode();
		L->right = Create_TreeNode();
	}
	return L;
}

//Print the node being visited now
void Visit(Tree T) {
	cout << T->val << " ";
}

void InOrder(Tree T) {
	if (T != NULL) {
		InOrder(T->left);
		Visit(T);
		InOrder(T->right);
	}
}

void InitList(List* p) {
	List q = *p;
	q->next = NULL;
	q->val = NULL;
}

void Insert(List* L, Tree T) { // Insert at head.
	List p = NEW(ListNode);
	List q = *L;
	p->next = q->next;
	q->next = p;
	p->val = T;
}

List Remove(List* T) { // Remove at head.
	List q = *T;
	List p = q->next;
	if (p != NULL) {
		q->next = p->next;
	}
	return p;
}

int Len(List T) {
	int i = 0;
	T = T->next;
	while (T != NULL) {
		T = T->next;
		i++;
	}
	return i;
}

void myprint(List L) {
	printf("len: %d [ ", Len(L));
	List p = L->next;
	while (p != NULL) {
		printf("%d ", p->val->val);
		p = p->next;
	}
	printf("]\n");
}

bool GetNodePath(Tree root, Tree T, List* L) {
	bool found = false;
	if (root == NULL)
		return false;
	Insert(L, root);
	if (root == T)
		return true;
	if (root->left != NULL)
		found = GetNodePath(root->left, T, L);
	if (!found && root->right != NULL)
		found = GetNodePath(root->right, T, L);
	if (!found)
		free(Remove(L));
	return found;
}

struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p,
		struct TreeNode* q) {
	if (root == NULL || p == NULL || q == NULL)
		return NULL;
	ListNode n1;
	ListNode n2;
	List a = &n1, b = &n2;
	InitList(&a);
	InitList(&b);
	GetNodePath(root, p, &a);
	GetNodePath(root, q, &b);
	myprint(a);     ///////////////
	myprint(b);     ///////////////
	List l1 = a->next, l2 = b->next;
	if (l1 == NULL || l2 == NULL)
		return NULL;
	int i;
	// Right align two lists.两个链表Y字型交叉找第一公共点问题
	if (Len(a) > Len(b)) {
		for (i = 0; i < (Len(a) - Len(b)); i++)
			l1 = l1->next;
	} else {
		for (i = 0; i < (Len(b) - Len(a)); i++)
			l2 = l2->next;
	}
	for (; l1->val != l2->val && l1 != NULL; l1 = l1->next, l2 = l2->next)
		;
	return l1 != NULL ? l1->val : NULL;
}

int main() {
	Tree T;
	T = Create_TreeNode();
	InOrder(T);
	// BinTree is: [1,2,#,#,3,#,#]
	cout << endl;
	cout << lowestCommonAncestor(T, T->left, T->right)->val << endl;
	return 0;
}
