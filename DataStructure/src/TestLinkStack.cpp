#include "LinkStack.h"
#include "stdio.h"

void testLinkStack(){
	printf("LinkStack test: \n");
	LinkStack L;
	int e = 0;
	InitStack_L(L);
	for(int i = 0; i < 10; i++)
		Push_L(L,i);
	GetTop_L(L,e);
	printf("%d\n", e);
	Pop_L(L,e);
	printf("%d\n",e);
	GetTop_L(L,e);
	printf("%d\n",e);
	while(!isEmpty_L(L)){
		Pop_L(L,e);
		printf("%d\n",e);
	}
	//GetTop_L(L,e);
	printf("%d",e);
	printf("\n");
}
