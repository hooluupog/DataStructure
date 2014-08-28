typedef int ElemType;
typedef int Status;

#ifndef _SQSTACK_H_
#define _SQSTACK_H_
#define OK 1
#define ERROR 0
#define EMPTY_STACK_TOP -1
#define STACK_SIZE 100
typedef struct Stack{
	ElemType top;  //栈顶游标
	ElemType data[STACK_SIZE]; //栈数据
}*SqStack;

Status InitStack_Sq(SqStack &s);//栈初始化
Status GetTop_Sq(SqStack s,ElemType &e);//返回栈顶元素到e中
Status Push_Sq(SqStack &s,ElemType e);//入栈
Status Pop_Sq(SqStack s,ElemType &e);//出栈
void printStack_Sq(SqStack s);
Status IsEmpty_Sq(SqStack s);
#endif
