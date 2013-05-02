#include "LinkStack.h"
#include "fatal.h"

void InitStack_L(LinkStack &L)
{
	L = (LinkStack)malloc(sizeof(LNode));
	if(L == NULL) Error("Init failure: out of space!");
	L->data = -1;
	L->next = NULL;
}
bool isEmpty_L(LinkStack L)
{
	return L == NULL;
}
Status GetTop_L(LinkStack &L,ElemType &e)
{
	if(L == NULL){
		Error("GetTop error:Empty Stack");
		return ERROR;//return value used to avoid warning
	}
	e = L->data;
	return OK;
}
Status Push_L(LinkStack &L,ElemType e)
{
	LinkStack p;
	p = (LinkStack)malloc(sizeof(LNode));
	if(!p){
		Error("Push failure:out of space !!");
		return ERROR;
	}
	p->data = e;
	p->next = L;
	L = p;
	return OK;
}
Status Pop_L(LinkStack &L,ElemType &e)
{
	LinkStack p;
	if(L == NULL){
		Error("Pop failure:Empty Stack");
		return ERROR;//return value used to avoid warning
	}
	p = L;
	e = L->data;
	L = p->next;  //删除栈顶元素
	free(p);  //释放已删除结点指针
	return OK;
}
