/* poj:1961 
 * 题目大意：给定字符串S，求其前n位重复的次数。比如aabaabaabaab，前2位是aa，a重复了2次，前6位是aabaab，aab重复了2次，
 *前9位是aabaabaab，aab重复了3次，前12位是aabaabaabaab，aab重复了4次。所以输出的结果就是
 * 2 2 
 * 6 2 
 * 9 3 
 * 12 4 
*/
#include "stdio.h"
#include "string.h"
#define MAX 1000010

char s[MAX];
int next[MAX],len;

void get_next()
{//next函数
  int i = 0,j = -1;
  next[0] = -1;
  while(i < len)
    {
      if(j == -1 || s[i] == s[j])
	{
	  i++;j++;
	  next[i] = j;
	}
      else
	j = next[j];
    }
}

int main()
{
  int n,i = 1,j;
  while(scanf("%d",&len) && len){
    scanf("%s",&s);
    get_next();
    printf("Test case #%d\n",i++);
    for(n = 2;n <= len; ++n)
      //next[n]表示最大前缀子串长度，当n%(n-next[n])==0时，
      //n-next[n]就表示最小前缀子串长度，并且每个最小子串之间是相邻的，
      //也就是说n/(n-next[n])表示该最小子串重复出现的次数。
      if((n%(j = n - next[n]) == 0) &&((n / j) > 1))
	printf("%d %d\n",n,n/j);
    printf("\n");
  }
  return 0;
}

