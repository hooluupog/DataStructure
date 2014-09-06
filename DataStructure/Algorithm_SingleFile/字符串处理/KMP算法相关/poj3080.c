/*给定m个串，求长度大等于3且在每个串中都出现的最长子串，如果两个子串
 * 长度一样要求输出字典序小的一个答案。如果没有则输出
 * “no significant commonalities".
 */

#include "stdio.h"
#include "string.h"
#define MAX 100

char ans[MAX],p[MAX],str[MAX][MAX];
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

int kmp_match(char* s, int sLen, char* p, int pLen,int* next) {
    int i = 0, j = 0;
    get_next(p, next, pLen);
    while (i < sLen && sLen >= pLen) {
        if (j == -1 || s[i] == p[j]) {
            i++;
            j++;
        } 
        else
            j = next[j];
        if (j == pLen) {
            return 1;
        }
    }
    return 0;
}

int main() {
    int n,m,i,j,k,pLen,len,flag;
    scanf("%d", &n);
    while(n--){
        scanf("%d", &m);
        for(i = 0;i <m;i++){
            scanf("%s",&str[i]);
        }
        ans[0] = '\0';
        flag = 0;
        pLen = strlen(str[0]);
        for(len = pLen;len>=3;len--){
            for(i = 0;i <=pLen-len;i++){
                for(k = 0,j = i;j <i+len;j++){
                    p[k++] = str[0][j];
                }
                p[k] = '\0';
                for(k=1;k<m;k++){
                    if(kmp_match(str[k],strlen(str[k]),p,strlen(p),next)==0) break;
                }
                if(k==m){
                    flag = 1;
                    if(strlen(ans)==strlen(p) && strcmp(ans,p) > 0){
                        strcpy(ans,p);
                    }  
                    if(strlen(ans) < strlen(p)){
                        strcpy(ans,p);
                    }
                }
            }
           if(flag == 1) break;
        }
        if(ans[0] == '\0') printf("no significant commonalities\n");
        else printf("%s\n",ans);
    }
    return 0;
}
