#include "stdio.h"
#include "memory.h"
#include <vector>
using namespace std;

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

//最小调整有序，找出索引m和n，只要将m和n之间的元素排好序，整个数组就是有序的。
//解法一：
int BoundIndex(vector<int> a, int start, int end, int val, int direction) {
	//direction: 0表示寻找左边界; 1表示寻找右边界.
	while (start <= end && val >= a[start])
		start++;
	if (direction == 1)
		start--;
	return start;
}

vector<int> findSegment1(vector<int> A, int n) {
	int left, right, i, j;
	vector<int> res;
	for (i = 1; i < n && A[i - 1] <= A[i]; i++)
		; //寻找无序区间左边界
	for (j = n - 2; j >= 0 && A[j] <= A[j + 1]; j--)
		; //无序区间右边界
	if (i == n) { //原数组已经有序
		left = 0;
		right = 0;
	} else {
		int min = A[i];
		int max = A[j];
		for (int k = n - 1; k >= i; k--) { //左边界以右的最小元素
			if (min > A[k])
				min = A[k];
		}
		for (int k = 0; k <= j; k++) { //右边界以左的最大元素
			if (max < A[k])
				max = A[k];
		}
		left = BoundIndex(A, 0, i, min, 0); //找最左左边界
		right = BoundIndex(A, j, n - 1, max, 1); //找最右右边界

	}
	res.push_back(left);
	res.push_back(right);
	return res;
}


//解法二
vector<int> findSegment2(vector<int> A, int n) {
	vector<int> res;
	int left, right;
	if (A.size() == 0 || n <= 0) {
		return res;
	}
	// 计算[left,right]中的right
	// 如果当前元素小于之前的最大元素则说明当前元素应处于[left,right]无序序列中
	// 并且当前元素是当前最大下标的无序元素，所以更新right为当前元素下标
	for (int i = 1, max = A[0]; i < n; i++) {
		if (max <= A[i]) {
			max = A[i];
		} else {
			right = i;
		}
	}
	// 计算[left,right]中的right
	// 如果当前元素大于之前的最小元素则说明当前元素应处于[left,right]无序序列中
	// 并且当前元素是当前最小下标的无序元素，所以更新left为当前元素下标
	for (int i = n - 2, min = A[n - 1]; i >= 0; i--) {
		if (min >= A[i]) {
			min = A[i];
		} else {
			left = i;
		}
	}
	res.push_back(left);
	res.push_back(right);
	return res;
}

//寻找最长递增子序列(动态规划)
//初始化max[n]所有元素为1；max[i]表示以a[i]为结尾的递增子序列长度
// max[0] = 1;
// max[i] = max[i]，if max[j] < max[i],(j = 0，1，2...，i-1)
// max[i] = max[k] + 1, max[k] = {max[0],max[1],...,max[i-1]},即[0,...,i-1]中最长的递增子序列长度。
// 最长的递增子序列长度： longest = max{max[0],max[1],...,max[n-1]}.
int findLongest(vector<int> A, int n) {
	if (A.size() == 0 || n <= 0)
		return 0;
	vector<int> max;
	int longest = 0;
	for (int i = 0; i < n; i++) {
		max.push_back(1);
		for (int j = i - 1; j >= 0; j--) {
			if (A[j] < A[i]) {
				max[i] = max[j] >= max[i] ? max[j] + 1 : max[i];
			}
		}
		longest = longest > max[i] ? longest : max[i];
	}
	return longest;
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

	int a[] = { 1, 2, 6, 5, 8, 9 };
	vector<int> v(a, a + sizeof(a) / sizeof(int));
	vector<int> res = findSegment1(v, 6);
	for (int i = 0; i < res.size(); i++) {
		printf("%d ", res[i]);
	}
	printf("\n");
	res = findSegment2(v, 6);
	for (int i = 0; i < res.size(); i++) {
		printf("%d ", res[i]);
	}
	printf("\n");
	int b[] = { 203, 39, 186, 207, 83, 80, 89, 237, 247 };
	vector<int> bv(b, b + sizeof(b) / sizeof(int));
	printf("%d\n", findLongest(bv, 9));

    return 0;
}
