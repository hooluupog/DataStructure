#include "LinkList.h"
#include "fatal.h"

Status ListInsert_L(LinkList &L,int i,ElemType e)
//在带头结点的单链表L的第i个元素之前插入元素e
{
	LinkList p,s;
	p = L;
	int j = 0;
	while(p && j<i-1)//寻找第i-1个结点
	{
		p = p->next;
		++j;
	}
	if(!p || j>i-1){
		Error("illegal insert index position");
	       	return ERROR;//i小于1或者大于表长
	}
	s = (LinkList)malloc(sizeof(LNode)); //生成新结点
	//插入到L中
	s->data = e;
	s->next = p->next;
	p->next = s;
	return OK;
}
Status ListDelete_L(LinkList &L,int i,ElemType &e)
//在带头结点的单链表L中，删除第i个元素，并由e返回其值
{
	LinkList p,q;
	p = L;
	int j = 0;
	while(p->next && j<i-1)//寻找第i个结点，并令p指向其前驱
	{
		p = p->next;
		++j;
	}
	if(!(p->next) || j > i-1) 
	{
		Error("illegal delete index position");
	       	return ERROR;//i值不合法
	}
	//删除并释放结点
	q = p->next;
	p->next = q->next;
	e = q->data;
	free(q);
	return OK;
}
void DestroyList_L(LinkList &L)
//销毁链表
{
	LinkList p;
	while(L)
	{
		p = L->next;
		free(L);
		L = p;
	}
}
Status GetElem_L(LinkList L,int i,ElemType &e)
//第i个元素存在时赋值给e并返回OK，否则返回ERROR
{
	LinkList p;
	p = L->next;
	int j = 1;
	while(p && j<i)
	{
		p = p->next;
		++j;
	}
	if(!p || j>i){
		Error("illegal index position");
	       	return ERROR; //第i个元素不存在
	}
	e = p->data; //取第i个元素
	return OK;
}
void CreateList_Head_L(LinkList &L, int n)
//头插法建立带头结点的单链表L，含有n个元素
{
	LinkList p;
	L = (LinkList)malloc(sizeof(LNode));
	L->next = NULL;
	for(int i=0; i<n; ++i)
	{
		 p = (LinkList)malloc(sizeof(LNode)); //生成新结点
		 scanf("%d",&p->data); //输入元素值
		 //插入到表头
		 p->next = L->next;
		 L->next = p;
	}
}
void CreateList_Tail_L(LinkList &L, int n)
//尾插法建立带头结点的单链表L，含有n个元素
{
	LinkList p,temp;
	L = (LinkList)malloc(sizeof(LNode));
	temp = L;//temp始终指向链表表尾
	L->next = NULL; //先建一个带头结点的单链表
	for(int i=0; i<n; ++i)
	{
		p = (LinkList)malloc(sizeof(LNode)); //生成新结点
		p->next = NULL;
		scanf("%d",&p->data); //输入元素值
		//插入到表尾
		temp->next = p;
		temp = p;
	}
}
