typedef int ElemType;
typedef int Status;

#ifndef _LINKQUEUE_H_
#define _LINKQUEUE_H_
#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE

typedef struct QNode{
	ElemType data;
	struct QNode *next;
}QNode,*QueuePtr;

typedef struct {
	QueuePtr front;//front pointer
	QueuePtr rear;//rear pointer
}LinkQueue;

Status Init_LinkQueue(LinkQueue &Q); //construct an empty queue
Status Destroy_LinkQueue(LinkQueue &Q); //destruct the queue
Status En_LinkQueue(LinkQueue &Q, ElemType e); //insert a new elem on rear
Status De_LinkQueue(LinkQueue &Q,ElemType &e); //delete elem on head assign it into e
bool IsEmpty_LinkQueue(LinkQueue Q);

#endif
