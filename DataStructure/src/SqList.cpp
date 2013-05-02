#include "SqList.h"
#include "fatal.h"

Status InitList_Sq(SqList &L)
{
//构造一个空的线性表
	L.elem = (ElemType *)malloc(LIST_INIT_SIZE*sizeof(ElemType));
	if(!L.elem) Error("Init failure: out of space!");
	L.length = 0;
	L.listSize = LIST_INIT_SIZE; //初始化存储容量
	return OK;
}

Status Create_SqList(SqList &L,int n)
{
	InitList_Sq(L);//初始化顺序表
	for (int i = 0; i < n; ++i) {//创建顺序表
		scanf("%d",&(L.elem[i]));
		L.length++;
		if(L.length >= L.listSize) { //当前存储空间已满，增加容量
				ElemType *newbase = (ElemType *)realloc(L.elem,(L.listSize+LISTINCREMENT)*sizeof(ElemType));
				if(!newbase){
					Error("realloc failure:out of space!");	
					return ERROR;//存储分配失败
				}
				L.elem = newbase; //新基址
				L.listSize += LISTINCREMENT;//增加存储容量
			}
	}
	return OK;
}//Create_SqList

Status ListInsert_Sq(SqList &L,int i,ElemType e)
{
//在顺序线性表L的第i个元素之前插入新元素e
//i的合法值为1<=i<=ListLength_Sq(L)+1
	ElemType *p,*q;
	if(i<1 || i>L.length+1){
		Error("illegal insert index position");
	       	return ERROR;//i值不合法
	}
	if(L.length >= L.listSize) { //当前存储空间已满，增加容量
		ElemType *newbase = (ElemType *)realloc(L.elem,(L.listSize+LISTINCREMENT)*sizeof(ElemType));
		if(!newbase){
			Error("realloc failure:out of space!");	
			return ERROR;//存储分配失败
		}
		L.elem = newbase; //新基址
		L.listSize += LISTINCREMENT;//增加存储容量
	}
	q = &(L.elem[i-1]);//q为插入位置
	for(p = &(L.elem[L.length-1]); p>=q; --p)
		*(p+1) = *p;			 //插入位置及之后的元素右移
		*q = e;                   //插入e
		++L.length;
		return OK;
}
Status ListDelete_Sq(SqList &L,int i,ElemType &e)
{
//在顺序线性表L中删除第i个元素，并用e返回其值
//i的合法值为1<=i<=ListLength_Sq(L)+1
	ElemType *p,*q;
	if(i<1||i>L.length) {
		Error("illegal delete index position");
	       	return ERROR;//i值不合法
	}
	p = &(L.elem[i-1]); //p为被删除元素的位置
	e = *p; //被删除元素的值赋给e
	q = L.elem + L.length-1; //表尾元素的位置
	for(++p; p<=q; ++p)
		*(p-1) = *p; //被删除元素之后的元素左移
	--L.length;
	return OK;

}
void DestroyList(SqList &L)
{
//释放顺序表L所占的存储空间
	free(L.elem);
	L.listSize = 0;
	L.length = 0;
}
int LocateElem_Sq(SqList L,ElemType e)
{
//在顺序表中查找第一个值为e的元素的位序
	ElemType *p = L.elem; //p初始化为第一个元素的位序
	int i = 1; 
	while(i<=L.length && (*p!=e))
	{
		++i;++p;
	}
	if(i<=L.length) return i;
	else return 0;
}
