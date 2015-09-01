#include "stdio.h"
#include "memory.h"

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

int main(){
    char dest[100];
    char src[] = "We are happy.";
    memcpy(dest,src,14);
    ReplaceBlank(dest,100);
    int i = 0;
    while(dest[i] != '\0'){
    	printf("%c",dest[i++]);
    }
    return 0;
}
