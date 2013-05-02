typedef int ElemType;
typedef int Status;

#ifndef _DULINKLIST_H_
#define _DULINKLIST_H_
#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE

typedef struct DuLNode{
	ElemType data;
	struct DuLNode *prior;
	struct DuLNode *next;
}DuLNode,*DuLinkList;

Status Create_DuLinkList(DuLinkList& DL, int n);//create a new DuLinkList
DuLinkList GetElemP_DuL(DuLinkList DL, ElemType e);// Get element
Status Insert_DuL(DuLinkList &DL,int i, ElemType e);//insert elem before index i
Status Delete_DuL(DuLinkList &DL,int i, ElemType e);//delete elem and assign its value into e 

#endif
