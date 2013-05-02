#include "LinkList.h"
#include "stdio.h"

//LinkList
void testLinkList()
{
	freopen("input/LinkList.txt","r",stdin);
	printf("LinkList test: \n");
//找到两个链表共同后缀的起始结点,2012考研计算机综合数据结构42题
	LinkList l1,l2,str1,str2,p;
	int mark = 0;//标记是否更新后缀起始结点位置
	//建立链表l1("1234567"),l2("89567")
	CreateList_Tail_L(l1,7);
	CreateList_Tail_L(l2,5);
	//让str1和str2分别指向头结点的下一个结点
	str1 = l1->next;
	str2 = l2->next;
	//长链表的指针后移两链表长度之差个位数
	for(int i = 0;i<2;++i)
		str1 = str1->next;
	//开始同步移动两个链表指针str1，str2
	while(str1 != NULL && str2 != NULL)
	{
		if(str1->data == str2->data && mark == 0)
		{
			p = str1;//p指向共同后缀起始结点
			mark = 1;//表示共同后缀结点位置发生变化
		}
		if(str1->data != str2->data)
			mark = 0;
		str1 = str1->next;
		str2 = str2->next;
	}
	if(mark == 1)
		printf("共同后缀的起始结点为：%d\n",p->data);
	else
		printf("没有共同后缀\n");
}
