/* 从一个大小为n的数组中找出前k大的元素 */
/* 如果数组彼此的元素都是不同的，那么也 */
/* 可以用哈希法去做                    */

#include "stdio.h"

int a[1000001];
void SiftDown(int* list,int root, int end) {
	int child;
	int siftdata;
	siftdata = list[root];
	for(child = 2*root; child <= end; child = 2*child) {
		if(child < end && list[child] < list[child + 1]) 
			child ++;
		if(siftdata >= list[child])
			break;
		list[root] = list[child];
		root = child;
	}
	list[root] = siftdata;
}

void HeapSort(int* list,int m) {
	int i,temp;
	for(i = list[0]/2; i >= 1; i--) {
		SiftDown(list,i,list[0]);
	}
	for(i = a[0]; i > a[0] - m; i--) {
		temp = list[i];
		list[i] = list[1];
		list[1] = temp;
		SiftDown(list,1,i-1);
	}
}

void main(){
	int n,m,i = 1;
	scanf("%d %d\n",&n,&m);
	a[0] = n;
	while(n--){
		scanf("%d",&a[i++]);
	}
	HeapSort(a,m);
	for(i = 1;i <= m-1;i++){
		printf("%d ",a[a[0] - i + 1]);
	}
	printf("%d\n",a[a[0] - i + 1]);
}
