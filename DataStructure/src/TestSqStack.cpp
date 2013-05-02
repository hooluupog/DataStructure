#include "stdio.h"
#include "SqStack.h"
/*汉诺塔递归版本
 * 将塔座x上按直径有小到大且自上而下编号为1至n的n个圆盘按规则搬到塔座z，y作为辅助
 * 盘，其中直径大的盘不能放在直径小的盘子上,每次只能移动一个盘
 * int count = 0; //全局变量，对搬动计数
 *	void hanoi(int n,char x,char y,char z)
 *	{
 *		if(n == 1)
 *			printf("%d. Move disk %d from %c to %c\n",++count,n,x,z);
 *		else{
 *			hanoi(n-1,x,z,y);
 *			printf("%d. Move disk %d from %c to %c\n",++count,n,x,z);
 *			hanoi(n-1,y,x,z);
 *		}
 *	}
*/

int count = 0; //全局变量，对搬动计数
void hanoi(int n,char x,char y,char z){
//栈模拟汉诺塔递归算法
	SqStack A,B,C,S,D;//栈A,B,C分别存储当前变量参数信息，
	//S存储n(当前需要移动的盘子数目)，D记录当前移动的是第几个盘子
	int xi,yi,zi,disk;
	xi = x;
	yi = y;
	zi = z;
	InitStack_Sq(A);
	InitStack_Sq(B);
	InitStack_Sq(C);
	InitStack_Sq(S);
	InitStack_Sq(D);
	Push_Sq(A,xi);
	Push_Sq(B,yi);
	Push_Sq(C,zi);
	Push_Sq(S,n);
	Push_Sq(D,n);
	//模拟递归开始
	while(!IsEmpty_Sq(A) && !IsEmpty_Sq(B) && !IsEmpty_Sq(C) && !IsEmpty_Sq(S) && !IsEmpty_Sq(D))
	{
		//pop出的值存入xi,yi,zi,n,disk，更新当前的变量参数信息，类似于递归程序中
		//函数参数传递，将参数信息传入子函数中
		Pop_Sq(A,xi);
		Pop_Sq(B,yi);
		Pop_Sq(C,zi);
		Pop_Sq(S,n);
		Pop_Sq(D,disk);
		if(n == 1) //n=1，模拟递归出口，disk记录要移动的盘号
			printf("%d. Move disk %d from %c to %c\n",++count,disk,xi,zi);
		else{
			//注意栈的压入顺序和递归函数的调用是正好相反的
			L1: //hanoi(n-1,y,x,z)
				Push_Sq(A,yi);
				Push_Sq(B,xi);
				Push_Sq(C,zi);
				Push_Sq(S,n-1);
				Push_Sq(D,n-1);
			L2: //move n  A->C
				Push_Sq(A,xi);
				Push_Sq(B,yi);
				Push_Sq(C,zi);
				Push_Sq(S,1);
				Push_Sq(D,n);
			L3: //hanoi(n-1,x,z,y)
				Push_Sq(A,xi);
				Push_Sq(B,zi);
				Push_Sq(C,yi);
				Push_Sq(S,n-1);
				Push_Sq(D,n-1);
		}
	}
}

  void testSqStack(){
	  printf("SqStack test: \n");
	  hanoi(3,'A','B','c');
  }
