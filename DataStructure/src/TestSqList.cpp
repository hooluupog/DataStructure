#include  "SqList.h"
#include "stdio.h"
#include "algorithm.h"

void invert(SqList L,int s,int t){
	//将顺序表中下标s至t的元素逆转
	int temp;
	for (int i = s; i <= (s+t)/2; ++i) {
		temp = L.elem[i];
		L.elem[i] = L.elem[t-i+s];
		L.elem[t-i+s] = temp;
	}
}//invert

void exchange(SqList L, int m,int n){
	//将顺序表中前m个元素和后n个元素交换位置(m+n=length)，采用“三次逆置”法,
	//如果m+n<length,即m和n之间还有元素存在，则“四次逆置”可得到结果。
	invert(L,0,m+n-1);
	invert(L,0,n-1);
	invert(L,n,m+n-1);
}
void testSqList() {
	freopen("input/SqList.txt","r",stdin);
	printf("SqList test: \n");
	ElemType e;
	SqList L,H;
	Create_SqList(L, 10);
	Create_SqList(H, 10);
	ListInsert_Sq(L, 5, 20);
	for (int i = 0; i < L.length; ++i)
		printf("%d ", L.elem[i]);
	printf("\n");
	ListDelete_Sq(L, 4, e);
	for (int i = 0; i < L.length; ++i)
		printf("%d ", L.elem[i]);
	int result = LocateElem_Sq(L, 9);
	printf("\n%d\n", result);
	exchange(L,3,L.length-3);
	for (int i = 0; i < L.length; ++i)
	{
		H.elem[i] = L.elem[i];
		printf("%d ", L.elem[i]);
	}
	printf("\n");
	Qsort(L,0,L.length-1);
	printf("Quick sort:\n");
	for (int i = 0; i < L.length; ++i)
		printf("%d ", L.elem[i]);
	printf("\n");
	HeapSort(H);
	printf("Heap sort:\n");
		for (int i = 0; i < H.length; ++i)
			printf("%d ", H.elem[i]);
		printf("\n");
}

