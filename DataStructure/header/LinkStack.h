typedef int ElemType;
typedef int Status;

#ifndef _LINKSTACK_h_
#define _LINKSTACK_h_
#define OK 1
#define ERROR 0
#define OWERFLOW -2
typedef struct LNode{
	ElemType data; //数据域
	struct LNode *next; //指针域
}*LinkStack;

void InitStack_L(LinkStack &L);
bool isEmpty_L(LinkStack L);
Status GetTop_L(LinkStack &L,ElemType &e);
Status Push_L(LinkStack &L,ElemType e);
Status Pop_L(LinkStack &L,ElemType &e);
#endif
