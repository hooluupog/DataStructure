// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
// 使得所有的奇数位于数组的前半部分，所有的偶数位于位于数组的
// 后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。
// 输入：
// 5
// 1 2 3 4 5
// 输出：
// 1 3 5 2 4

#include "stdio.h"
#include "stdlib.h"

typedef struct LNode {
    int data;
    LNode *next;
    LNode *pre;
}*LinkList;

void OddEven(LinkList &L){
    LinkList p,s,q;
    s = L; // s记录扫描时的奇数位
    p = L->next; // 工作指针
    while(p != NULL){
        q = p -> next; //保存p的下一个位置
        if ((p ->data & 1) !=0){// 当前位置为奇数
            if(s->next != p) { // 两个奇数位不相邻
                // 将奇数插入到上一个奇数的后继位置
                p->pre->next = p->next; 
                s->next->pre = p;
                if (p->next != NULL){
                    p->next->pre = p->pre;
                }
                p->pre = s;
                p->next = s->next;
                s->next = p;
            }
            //更新最新扫描过的奇数的位置
            s = s->next;
        }
        // p指向下一位准备下一次循环
        p = q;
    }
}

int main() {
    int n,i;
    LinkList L,S,p;
    L = (LinkList)malloc(sizeof(LNode));
    L->next = NULL;
    L->pre = NULL;
    p = L;
    scanf("%d",&n);
    for (i = 0; i < n; i++) {
        S = (LinkList)malloc(sizeof(LNode));
        scanf("%d",&S->data);
        //后插法创建链表
        p->next = S; 
        S->pre = p; 
        p = S;
        S->next = NULL;
    }
    OddEven(L);
    for (i = 0; i < n-1; i++) {
        L = L->next;
        printf("%d ",L->data);
    }
    L = L->next;
    printf("%d\n",L->data);
    return 0;
}
