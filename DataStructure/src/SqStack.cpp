#include "SqStack.h"
#include "fatal.h"

Status InitStack_Sq(SqStack &s)
//栈初始化
{
	s = (SqStack)malloc(sizeof(Stack));
	if(!s) Error("Init failure: out of space!");
	s->top = EMPTY_STACK_TOP;
	return OK;
}
Status GetTop_Sq(SqStack s,ElemType &e)
//返回栈顶元素到e中
{
	if(s->top == EMPTY_STACK_TOP)
	{
		Error("GetTop error:Empty Stack");
		return ERROR;//return value used to avoid warning
	}
	else{
		e = s->data[s->top];
		return OK;
	}
}
Status Push_Sq(SqStack s,ElemType e)
//入栈
{
	if(s->top == STACK_SIZE-1)
	{
		Error("Push failure:out of space !!");
		return ERROR;
	}
	else s->data[++s->top] = e; //插入新的元素
		return OK;
}
Status Pop_Sq(SqStack s,ElemType &e) 
//出栈
{
	if(s->top == -1)
	{
		Error("Pop failure:Empty Stack");
		return ERROR;//return value used to avoid warning
	}
	else e = s->data[s->top--];
		return OK;
}

void printStack_Sq(SqStack s)
{
	int top = s->top; //保存原始top值
	while(s->top != -1) printf("%3d",s->data[s->top--]);
	printf("\n");
	s->top = top;
}
Status IsEmpty_Sq(SqStack s)
{
	if(s->top == -1) return OK;
	else return ERROR;
}	
