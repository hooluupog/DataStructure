/*
 *   给出字符串找出所有的前缀和后缀相等的子串
 *   按小到大输出这些子串的长度
 *   输入:
 *   ababcababababcabab
 *   aaaaa
 *
 *
*/
#include "stdio.h"
#include "string.h"
#define MAX 400005
char s[MAX];
int next[MAX],answer[MAX],length;

void get_next()
{
  int i = 0, j = -1;
  next[0] = -1;
  while(i < length)
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
  freopen("data.txt","r",stdin);
  while(scanf("%s",&s) != EOF)
    {
      length = strlen(s);
      get_next();
	  answer[0] = length;
      int i = length,n = 0;
      while(next[i] > 0)
	  {
	    answer[++n] = next[i];
		i = next[i];
	  }
      for(i = n; i >= 0; --i)
		printf("%d ",answer[i]);
      printf("\n");
    }
  return 0;
}
