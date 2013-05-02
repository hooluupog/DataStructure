/*基于顺序存储结构的循环队列(带头结点)*/
#include "SqQueue.h"
#include "fatal.h"

//Construct an empty queue
Status Init_SqQueue(SqQueue &Q){
	Q.data = (ElemType*)malloc(MAXQSIZE*sizeof(ElemType));
	if(!Q.data){
		Error("out of space");
		return ERROR;
	}
	Q.front = Q.rear = 0;
	return OK;
}//Init_SqQueue

//Get current length of the queue
Status Length_SqQueue(SqQueue &Q){
	return (Q.rear - Q.front + MAXQSIZE) % MAXQSIZE;//取模
}//Length_SqQueue

 //入队
Status En_SqQueue(SqQueue &Q, ElemType e) {
	if((Q.rear + 1) % MAXQSIZE == Q.front){//队满判断,队满时front和rear是相邻的
		Error("Queue already full");
		return ERROR;
	}
	Q.data[Q.rear] = e;
	Q.rear = (Q.rear + 1) % MAXQSIZE;//队尾指针后移
	return OK;
}//En_SqQueue


//出队,用e保存出队元素值
Status De_SqQueue(SqQueue &Q,ElemType &e){
	if(Q.rear == Q.front) {
		Error("Queue already empty");
		return ERROR;
	}
	e = Q.data[Q.front];
	Q.front = (Q.front + 1) % MAXQSIZE; //队头指针后移
	return OK;
}//De_SqQueue
