/*顺序统计量问题，求数组中第i小的元素*/

/*见算法导论中位数那章中select算法部分*/

#include "stdio.h"

int partition(int A[],int p,int r)
{//划分数组，使得左边的子数组元素都小于pivot（中轴），右边的子数组的元素大于等于pivot，并返回pivot的下标
  int pivot,i,j,temp;
  pivot = A[r];//取最后一个元素为中轴
  i = p-1;
  for(j = p;j <= r-1;++j)
    {
      if(A[j] <= pivot)
	{
	  i += 1;
	  temp = A[i];
	  A[i] = A[j];
	  A[j] = temp;
	}
    }
  //确定中轴元素所在的位置
  A[r] = A[i+1];
  A[i+1] = pivot;
  return i+1;
}

int select(int A[],int p,int r,int i)
{//找出数组中第i小的元素
  int q,k;
  if(p == r)
    return A[r];//数组只有一个元素
  q = partition(A,p,r);//得到划分后中轴元素的下标
  k = q - p + 1;
  if(i==k)
    return A[q];//中轴元素就是所要的元素
  else if(i < k)
    return select(A,p,q-1,i);//在左子数组中查找
  else
    return select(A,q+1,r,i-k);//在右子数组中查找
}

int main()
{
  int A[20] = {2,1,3,8,9,12,45,23,45,32,8,67,19,43,53,48,6,20,0,30};
  int result = select(A,0,19,10);
  printf("%d",result);
  return 0;
}
