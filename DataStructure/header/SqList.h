typedef int ElemType;
typedef int Status;

#ifndef _SQLIST_H_
#define _SQLIST_H_
#define OK 1
#define ERROR 0
#define LIST_INIT_SIZE 100
#define LISTINCREMENT 10
struct SqList{
ElemType *elem;//存储空间基址
int length;//表长
int listSize;//当前分配的存储容量(以sizeof(ELemType)为单位)
};

/*注意：如果把struct的定义放在cpp文件中，h文件中只放struct的声明，则需要：
**typedef struct SqList *ListPtr;
**typedef ListPtr List;
**接下来，在其他的cpp文件中声明时，使用List，如List L;
**如果使用SqList L，则会出现“struct未定义”的错误。
**因为：#include预编译指令是把.h文件的内容地当作一段文本嵌入到.c内，
**一个.c文件（对编译器来说就是一个可独立编译的单元）内不能存在未定义的类型，
**因为如果存在未定义类型时，编译器独立编译它的时候不知道要为该类型分配多大内存空间。
**但是允许有未定义的函数，因为对于函数调用而言编译器仅仅需要知道一个函数地址就可以了，
**而定义可以到连接时再作，所以就允许编译器有未定义的函数，但是连接时如果存在未定义的函数就会报错！！
**所以，声明一个指向结构体的指针，其他cpp文件引用List时只用到了结构体的入口地址，要访问内部的数据成员时通过函数调用即可。
*/

Status InitList_Sq(SqList &L);//初始化一个空的线性表
Status Create_SqList(SqList &L,int n);//创建一个线性表
Status ListInsert_Sq(SqList &L,int i,ElemType e);//在顺序线性表L的第i个元素之前插入新元素e
Status ListDelete_Sq(SqList &L,int i,ElemType &e);//在顺序线性表L中删除第i个元素，并用e返回其值
void DestroyList_Sq(SqList &L);//释放顺序表L所占的存储空间
int LocateElem_Sq(SqList L,ElemType e);//在顺序表中查找第一个值为e的元素的位序
#endif
