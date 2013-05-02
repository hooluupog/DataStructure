typedef int ElemType;
typedef int Status;

#ifndef _SQQUEUE_H_
#define _SQQUEUE_H_
#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE

#define MAXQSIZE 100
typedef struct{
	ElemType *data;
	int front;//若队列非空,指向队列头元素
	int rear;//若队列非空,指向队尾元素的下一个位置
}SqQueue;

Status Init_SqQueue(SqQueue &Q); //Construct an empty queue
Status Length_SqQueue(SqQueue &Q); //Get current length of the queue
Status En_SqQueue(SqQueue &Q, ElemType e); //入队
Status De_SqQueue(SqQueue &Q,ElemType &e); //出队,用e保存出队元素值

#endif
