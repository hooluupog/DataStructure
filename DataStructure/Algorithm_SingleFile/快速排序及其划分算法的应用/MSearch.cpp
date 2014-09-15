/*求中位数算法，划分算法的应用*/
#include "stdio.h"
 int M_Search(int a[], int b[],int n) 
 {
 	int s1=0,d1=n-1,m1,s2=0,d2=n-1,m2;
 	while(s1 != d1 || s2 != d2)
 	{
 		m1 = (s1+d1)/2;
 		m2 = (s2+d2)/2;
 		if (a[m1] == b[m2])
			return a[m1];
		if(a[m1]<b[m2])
		{
			if((s1+d1)%2 ==0)
			{
				s1 = m1;
				d2 = m2;
			}
			else
			{
				s1 = m1+1;
				d2 = m2;
			}
		}
		else
		{
			if((s1+d1)%2 ==0)
			{
				s2 = m2;
				d1 = m1;
			}
			else
			{
				s2 = m2+1;
				d1 = m1;
			}
		}
 	}
 	return a[s1] < b[s2] ? a[s1] : b[s2];
 }
  
 int main(int argc, char const *argv[])
 {
 	int a[20],b[20],result,i;
 	for (i = 0; i < 20; ++i)
 	{
 		a[i] = 4*i;
 		b[i] = 2*(i+1)+3;
		printf("%4d%4d\n",a[i],b[i]);
 	}
 	result = M_Search(a,b,20);
 	printf("%d\n",result);
 	return 0;
 }
