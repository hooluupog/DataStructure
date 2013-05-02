typedef int ElemType;
typedef int Status;

#ifndef _LINKLIST_H_
#define _LINKLIST_H_
#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE
typedef struct LNode{
	ElemType data;
	struct LNode *next;
}*LinkList;

Status ListInsert_L(LinkList &L,int i,ElemType e);//在带头结点的单链表L的第i个元素之前插入元素e
Status ListDelete_L(LinkList &L,int i,ElemType &e);//在带头结点的单链表L中，删除第i个元素，并由e返回其值
void DestroyList_L(LinkList &L);//销毁链表
Status GetElem_L(LinkList L,int i,ElemType &e);//第i个元素存在时赋值给e并返回OK，否则返回ERROR
void CreateList_Head_L(LinkList &L, int n);//头插法建立带头结点的单链表L，含有n个元素
void CreateList_Tail_L(LinkList &L, int n);//尾插法建立带头结点的单链表L，含有n个元素
#endif
