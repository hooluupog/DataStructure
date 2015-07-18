#include "stdio.h"
#define SIZE 1001
int pre[SIZE],in[SIZE];
void postOrder(int pstart,int istart,int n,int isRoot){
	int i;
	if (n < 1) return;
	for(i=0;pre[pstart] != in[istart+i];i++);//find the position of root
	// node which divides tree into two subtrees in inorder list.
	// traverse the left and right subtrees.
	postOrder(pstart+1,istart,i,0);
	postOrder(pstart+i+1,istart+i+1,n-i-1,0);
	isRoot? printf("%d",pre[pstart]):printf("%d ",pre[pstart]);
}
void testPostOrder(){
	int n,i;
	while(scanf("%d",&n) != EOF){
		for(i = 0;i<n;i++)
			scanf("%d",&pre[i]);
		for(i = 0;i<n;i++)
			scanf("%d",&in[i]);
		postOrder(0,0,n,1);
		printf("\n");
	}
}
