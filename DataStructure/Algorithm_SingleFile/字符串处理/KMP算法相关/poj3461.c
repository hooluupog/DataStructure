/*找公共子串*/
#include "stdio.h"
#include "string.h"
#define MAX 1000000

char s[MAX], p[MAX];
int next[MAX];
void get_next(char* s, int* next, int len) {
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

void match(char* s, int sLen, char* p, int pLen,int* next) {
    int i = 0, j = 0, count = 0;
    get_next(p, next, pLen);
    while (i < sLen && sLen >= pLen) {
        if (j == -1 || s[i] == p[j]) {
            i++;
            j++;
        } 
        else
            j = next[j];
        if (j == pLen) {
            count++;
            j = next[j]; 
        }
    }
    printf("%d\n",count);
}

int main() {
    int n,i;
    scanf("%d", &n);
    for(i = 0;i<n;i++){
        scanf("%s%s", &s, &p);
        strlen(s) > strlen(p)? match(s, strlen(s), p, strlen(p),next):match(p, strlen(p), s, strlen(s),next);
    }
    return 0;
}
