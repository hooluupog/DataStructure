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
 * 求二叉树的镜像；
 * 判断一棵树是否为另一颗树的子结构，如下图，B为A的子结构
 *          1                        8
 *        /   \                     /  \
 *       8     7                   9    2
 *     /   \
 *    9     2
 *         /  \
 *        4    7
 *        (A)                        (B)
 *
 * 判断是否平衡2叉树
 * 最近的公共祖先结点
 * 寻找中序遍历的下一个结点
 */

#include "stdio.h"
#include "BinTree.h"
#include "stdlib.h"
#include <vector>

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

//输出树的镜像
void MirrorTree(BinTree r){
	if (r == NULL) return;
	BinNode *temp;
	temp = r->lchild;
	r->lchild = r->rchild;
	r->rchild = temp;
	MirrorTree(r->lchild);
	MirrorTree(r->rchild);
}

bool doesHaveAllNodes(BinTree a,BinTree b){ // 以N为根结点遍历a，判断a是否含有和b相同结构的结点
	if(b == NULL) return true;  //到达b的叶子结点，说明具有相同结构，终止递归；
	if(a == NULL) return false; //到达a的叶子结点，说明不具有相同结构，终止递归；
	if(a ->data != b->data) return false; //结点值不相同，不具有相同结构，终止递归；
	return doesHaveAllNodes(a->lchild,b->lchild) && doesHaveAllNodes(a->rchild,b->rchild);
}
bool hasSubTreeRoot(BinTree a,BinTree b){ //判断b的根结点是否存在于a中，如果存在设此结点为N
	bool result = false;
	if (a->data == b->data)
		result = doesHaveAllNodes(a, b);
	if (!result && a->lchild != NULL)
		result = hasSubTreeRoot(a->lchild, b);
	if (!result && a->lchild != NULL)
		result = hasSubTreeRoot(a->rchild, b);
	return result;
}
bool IsSubTree(BinTree a,BinTree b){ //判断b是否为a的子结构
	if (a == NULL && b == NULL) return true;
	if ((a == NULL && b != NULL) || (a != NULL && b == NULL)) return false;
	return hasSubTreeRoot(a,b);
}

//找出根结点到所要查询的结点的路径存储到链表中
bool GetNodePath(BinTree r,int v,std::vector<BinNode*> &list){
	if (ToInt(r) == v) {
		return true;
	}
	list.push_back(r);
	bool found = false;
	if(r->lchild != NULL){
		found = GetNodePath(r->lchild,v,list);
	}
	if(!found && r->rchild != NULL){
	    found = GetNodePath(r->rchild,v,list);
	}
	if (!found){
		list.pop_back();
	}
	return found;
}

//找出二叉树中两个结点的最近公共祖先
BinNode *LastCommonParent(BinTree r, int v1, int v2) {
	BinNode *pLast = NULL;   // 最近的公共祖先结点
	std::vector<BinNode*> list1, list2; // 根结点到v1(v2)结点的路径链表
	GetNodePath(r, v1, list1);
	GetNodePath(r, v2, list2);
	//问题转化为求两个链表的首个公共结点
	if (list1.empty() || list2.empty())
		return NULL;
	std::vector<BinNode*>::iterator iter1, iter2;
	iter1 = list1.begin();
	iter2 = list2.begin();
	while (iter1 != list1.end() && iter2 != list2.end()) {
		if (*iter1 == *iter2) {
			pLast = *iter1;
		}
		iter1++;
		iter2++;
	}
	return pLast;
}

bool find(BinTree root, int node[], int p) {
	bool found = false;
	if (root->lchild != NULL) {
		found = find(root->lchild, node, p);
	}
	if (!found) {
		if (ToInt(root) == p) {
			node[0] = p;
		} else {
			if (node[0] == p) {
				node[1] = ToInt(root);
				return true;
			}
		}
		if (root->rchild != NULL) {
			found = find(root->rchild, node, p);
		}
	}
	return found;
}
//寻找中序遍历的下一个结点
int findSucc(BinTree root,int p){
	// write code here
	if (root == NULL)
		return -1;
	int node[2]; //存储p结点和p的后继结点
	bool found = false;
	found = find(root, node, p);
	if (found){
		return node[1];
	}
	return -1;
}

void testProblems() {
	BinTree r, s;
	freopen("input/BinTree1.txt", "r", stdin);
	r = Create_BinTree();
	s = Create_BinTree();
	printf("%d\n", IsSubTree(r, s));
	freopen("input/BinTree2.txt", "r", stdin);
	int res = 0;
	r = Create_BinTree();
	MaxPathSum(r, res);
	printf("%d\n", res);
	printf("%d\n", MinDepth(r));
	printf("%d\n", MaxDepth(r));
	printf("%d\n", IsAvlTree(r));
	printf("%d\n", IsSymmetricTree(r));
	freopen("input/BinTree4.txt", "r", stdin);
	r = Create_BinTree();
	PreOrder(r);
	printf("\n");
	MirrorTree(r);
	PreOrder(r);
	printf("\n");
	freopen("input/BinTree3.txt", "r", stdin);
	r = Create_BinTree();
	BinNode *p = LastCommonParent(r, 10, 13);
	if (p) {
		printf("%d\n", ToInt(p));
	}
	printf("\n");
    int node = findSucc(r,6);
    printf("%d\n", node);
}
