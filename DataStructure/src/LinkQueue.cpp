/* LinkQueue.cpp, 带头结点*/

#include "LinkQueue.h"
#include "fatal.h"

//construct an empty queue
Status Init_LinkQueue(LinkQueue &Q) {
	Q.front = Q.rear = (QueuePtr) malloc(sizeof(QNode));
	if (!Q.front) {
		Error("out of space");
		return ERROR;
	}
	Q.front->next = NULL;
	return OK;
} //InitQueue

//destruct the queue
Status Destroy_LinkQueue(LinkQueue &Q) {
	while (Q.front) {
		Q.rear = Q.front->next; //尾结点指向头结点的下一个位置
		free(Q.front); //释放头结点指针所指向的空间
		Q.front = Q.rear; //头结点指向尾结点,最终头结点和尾结点都为空
	}
	return OK;
}

//insert a new elem on rear
Status En_LinkQueue(LinkQueue &Q, ElemType e) {
	QueuePtr p;
	p = (QueuePtr)malloc(sizeof(QNode));
	if (!p) {
		Error("out of space");
		return ERROR;
	}
	p->data = e;
	p->next = NULL;
	Q.rear->next = p;
	Q.rear = p; //将新的结点作为尾结点
	return OK;
} //EnQueue

//delete elem on head assign it into e
Status De_LinkQueue(LinkQueue &Q, ElemType &e) {
	QueuePtr p;
	if (Q.front == Q.rear) {
		Error("empty Queue");
		return ERROR;
	}
	p = Q.front->next;
	e = p->data;
	Q.front->next = p->next;
	if (Q.rear == p)
		Q.rear = Q.front; //删除唯一存在的元素后,队尾指针悬空,需要给队尾指针重新赋值(指向头接点).
	free(p);
	return OK;
} //Dequeue

bool IsEmpty_LinkQueue(LinkQueue Q){
	return Q.front == Q.rear;
}
