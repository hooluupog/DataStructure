/*反转字符串
 *输入hello world
 *输出olleh dlrow
*/
#include "stdio.h"
#include "ctype.h"
#include "string.h"

void Reverse(char* p,int n)
{
    if(--n > 0) Reverse(p+1,n);
    printf("%c",*p);
    fflush(stdout);
}
int main()
{
  //setbuf(stdin,NULL);
  //setbuf(stdout,NULL);
  // setvbuf(stdin,NULL,_IONBF,0);
  // setvbuf(stdout,NULL,_IONBF,0);
  //freopen("input.txt","r",stdin);
    char s[255],*p,*q;
    int i=0,n;
    fgets(s,255,stdin);
    s[strlen(s)-1] = ' ';
    p = q = s;
    while(*p != '\0')
    {
        i++;
        p++;
        if(isspace(*p))
        {
            n = i;
            i = 0;
            Reverse(q,n);
            q = p+1;
        }
    }
    n = i;
    Reverse(q,n);
    printf("\n");
    return 0;
}
