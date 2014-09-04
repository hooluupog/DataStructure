/*kmp算法*/
#include "stdio.h"
#include "stdlib.h"
#include "string.h"

void get_next(char* s, int* next, int len) { //next函数
	int i = 1, j = -1;
	next[0] = -1;
	while (i < len) {
		if (j == -1 || s[i] == s[j]) {
			j++;
			next[i] = j;
			i++;
		} else
			j = next[j];
	}
}

void kmp_match(char* s, int sLen, char* p, int pLen) {
	int i = 0, j = -1, mark = 0, k = 0;
	//多申请一个int大小的空间，用于编译器存放数组大小信息，可以防止free函数执行时报堆溢出错误
	int* next = (int*) malloc(sizeof(int) * (pLen + 1));
	get_next(p, next, pLen);
	for (; i < sLen && sLen >= pLen; ++i) {
		while (j < pLen) {
			if (j == -1 || s[i] == p[j]) {
				j++;
				break;
			} //继续比较后继字符
			else
				j = next[j]; //模式串向右滑动
		}
		if (j == pLen) {
			printf("移动%d位，发生第%d次匹配\n", i - pLen + 1, ++k); //匹配成功
			mark = 1;
			j = next[j - 1]; //寻找下一次匹配
		}
	}
	if (mark == 0)
		printf("匹配失败\n");
	free(next);
}

int main() {
	char s[100], p[100];
	scanf("%s%s", &s, &p);
	kmp_match(s, strlen(s), p, strlen(p));
	return 0;
}
