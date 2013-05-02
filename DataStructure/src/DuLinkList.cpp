#include "DuLinkList.h"
#include "fatal.h"

//双向循环链表

//create a new DuLinkList
Status Create_DuLinkList(DuLinkList& DL, int n) {
	DuLinkList p,temp;
	DL = (DuLinkList)malloc(sizeof(DuLNode));
	if(DL == NULL) {
	       	Error("create DuLinkList fails:out of space");
		return ERROR;
	}
	DL->prior = NULL;
	DL->next = NULL;
	temp = DL;
	for(int i = 0; i < n; i++) {
		p = (DuLinkList)malloc(sizeof(DuLNode));
		if(p == NULL) {
			Error("Receive DuLinkList values fails:out of space");
			return ERROR;
		}
		scanf("%d",&p->data);
		temp->next = p; //temp is p prior node 
		p->prior = temp; // p->prior point to temp
		p->next = NULL;// p->next point to NULL
		temp = p;// temp move to p node
	}
	//set header prior pointer and tailer next pointer
	temp->next = DL;
	DL->prior = temp;
	return OK;
}//Create_DuLinkList

//Get element
DuLinkList GetElemP_DuL(DuLinkList DL, ElemType e) {
	// DL is a head pointer of DuLinkList with a head node
	DuLinkList p;
	p = DL->next; // Initialize: p point to the first node
	while(p != DL && p->data != e) {//clockwise search
		p = p->next;
	}
	if (p == DL) {
		Error("No such a element found");
		return ERROR;
	}
	else return p;
}//GetElemP_DuL

//insert elem before index i
Status Insert_DuL(DuLinkList &DL,int i, ElemType e) {
	DuLinkList p,s;
	p = DL;
	for(int j = 0; j <= i && p != NULL; j++)
	       	p = p->next;
	if (p == NULL) {
		Error("index is illegal");
		return ERROR;
	}
	if(!(s = (DuLinkList)malloc(sizeof(DuLNode))))
		Error("out of space");
	return ERROR;
	s->data = e;
	s->next = p;
	s->prior = p->prior;
	p->prior->next = s;
	p->prior = s;
	return OK;
}//Insert_DuL

//delete elem and assign its value into e 
Status Delete_DuL(DuLinkList &DL,int i, ElemType e) {
	DuLinkList p;
	p = DL;
	for(int j = 0; j <= i && p != NULL; j++)
	       	p = p->next;
	if (p == NULL) {
		Error("index is illegal");
		return ERROR;
	}
	p->prior->next = p->next;
	p->next->prior = p->prior;
	free(p);
	return OK;
} //Delete_DuL
