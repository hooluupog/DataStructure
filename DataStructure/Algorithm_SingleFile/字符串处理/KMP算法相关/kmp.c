/*kmp算法*/
#include "stdio.h"
#include "string.h"

void get_next(char* s, int* next, int len) { //next函数
    int i = 0, j = -1;
    next[0] = -1;
    while (i < len) {
        if (j == -1 || s[i] == s[j]) {
            i++;
            j++;
            next[i] = j;
        } else
            j = next[j];
    }
}

void kmp_match(char* s, int sLen, char* p, int pLen) {
    int i = 0, j = 0, mark = 0, k = 0;
    int* next = (int*) malloc(sizeof(int) * (pLen+1));
    get_next(p, next, pLen);
    while (i < sLen && sLen >= pLen) {
        if (j == -1 || s[i] == p[j]) {
            i++;
            j++;
        } //继续比较后继字符
        else
            j = next[j]; //模式串向右滑动
        if (j == pLen) {
            printf("移动%d位，发生第%d次匹配\n", i - pLen, ++k); //匹配成功
            mark = 1;
            j = next[j]; //寻找下一次匹配
        }
    }
    if (mark == 0)
        printf("匹配失败\n");
    free(next);
}

int main() {
    char s[100], p[100];
    while(scanf("%s%s", &s, &p) != EOF){
        kmp_match(s, strlen(s), p, strlen(p));
    }
    return 0;
}
