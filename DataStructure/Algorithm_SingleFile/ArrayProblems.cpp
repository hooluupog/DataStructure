#include "stdio.h"
#include "memory.h"

// 用'20%'替换字符串中的空格。
void ReplaceBlank(char string[], int length){
	//length为字符数组总容量，假设string[]空间足够大
   if(string == NULL && length <= 0) return;
   int originlength = 0; //原始字符串长度
   int numberofblank = 0;
   int i = 0;
   while(string[i++] != '\0'){
	   originlength++;
	   if(string[i] == ' '){
		   numberofblank++;
	   }
   }
   //空格替换为“20%”之后的字符串长度
   int newlength = originlength + numberofblank*2;
   int oIndex = originlength; //oIndex指向原始字符数组末尾
   int nIndex = newlength; //nIndex指向新的字符数组末尾
   while(oIndex >= 0 && oIndex < nIndex){ // oIndex==nIndex时，移动替换完毕
	   if(string[oIndex] != ' '){ //复制字符串
		   string[nIndex--] = string[oIndex];
	   }else{
		   // 替换空格
		   string[nIndex--] = '0';
		   string[nIndex--] = '2';
		   string[nIndex--] = '%';
	   }
	   oIndex--;
   }
}


unsigned int FirstBitIs1(int num){
    int indexBit = 0;
    while(((num & 1) == 0) && (indexBit < 8*sizeof(int))){
    	num  = num >> 1;
    	indexBit++;
    }
    return indexBit;
}

bool IsBit1(int num,int indexof1){
	return (num >> indexof1) & 1;
}

// 寻找整型数组中仅出现过一次的两个不同数(其他的数都出现过两次)。
void AppearOnce(int data[],int *num1,int *num2){
	if (data == NULL || data[0] < 2) return;
	int ExclusiveOR =  0,i;
	for(i=1;i<=data[0];i++){// a[0]为哨兵位
		ExclusiveOR ^= data[i]; //异或结果
	}
	unsigned int indexof1 = FirstBitIs1(ExclusiveOR);  //第一个非零位
	*num1 = *num2 = 0;
	for(i = 1;i <= data[0];i++){
		if(IsBit1(data[i],indexof1)){ //对应位非零的数在一起相异或
			*num1 ^= data[i];
		}else{ // //对应位为零的数在一起相异或
			*num2 ^= data[i];
		}
	}// 最终*num1和*num2为所求的两个不同数
}

int GetFirstK(int data[],int low,int high,int k){
    // 找出最左边k的位置
    if (high < low) return -1;
    int mid = (low + high) >> 1;
    if (data[mid] == k){
        if (mid > 0 && data[mid-1] == k){
            high = mid - 1;
        }else{
            return mid;
        }
    }else if(data[mid] > k){
        high = mid - 1;
    }else{
        low = mid + 1;
    }
    return GetFirstK(data,low,high,k);
}

int GetLastK(int data[],int low,int high,int k){
    // 找出最右边k的位置
    if (high < low) return -1;
    int mid = (low + high) >> 1;
    if (data[mid] == k){
        if (mid+1 <= high && data[mid+1] == k){
            low = mid + 1;
        }else{
            return mid;
        }
    }else if(data[mid] > k){
        high = mid - 1;
    }else{
        low = mid + 1;
    }
    return GetLastK(data,low,high,k);
}

// 有序数组中某个元素出现的次数
int GetNumOfK(int data[],int length,int k){
    if (length <= 0 || data == NULL) return 0;
    int right =  GetLastK(data,0,length-1,k);
    int left =   GetFirstK(data,0,length-1,k);
    if (right > -1 && left > -1){
        return right - left + 1;
    }
    return 0;
}

int main(){
	char dest[100];
	char src[] = "We are happy.";
	memcpy(dest, src, 14);
	ReplaceBlank(dest, 100);
	int i = 0;
	while (dest[i] != '\0') {
		printf("%c", dest[i++]);
	}
	int a[] = { 8, 2, 4, 3, 6, 3, 2, 5, 5 };
	int num1, num2;
	AppearOnce(a, &num1, &num2);
	printf("%d %d\n", num1, num2);
	int b[] = {1,2,3,3,3,3,4,5};
	printf("%d\n",GetNumOfK(b,8,3));
    return 0;
}
