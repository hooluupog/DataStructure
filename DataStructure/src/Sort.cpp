#include "algorithm.h"
#include "fatal.h"

void Qsort(UnsortedType &list,int low,int high) { // quicksort
    if(low < high){//表长度大于1
        int i,j,povit;
        i = low;
        j = high;
        povit = list.elem[i]; //取左端元素为中轴元素
        while(i < j){//i==j时，扫描完毕
            while(list.elem[j] > povit && i < j) j--;//从i位置向前扫描找小于等于中轴的元素
            if(i < j){
                list.elem[i++] = list.elem[j];
            }
            while(list.elem[i] < povit && i < j) i++;//从j位置向后扫描找小于等于中轴的元素
            if(i < j){
                list.elem[j--] = list.elem[i];
            }
        }//while
        list.elem[i] = povit;//一趟走完后中轴元素的位置
        Qsort(list,low,i-1);//快排中轴左边的元素
        Qsort(list,i+1,high);//快排中轴右边的元素
    }//if

}

void SiftDown(UnsortedType &list,int root, int end) { // heap adjust
    //List[start + 1,...,end] is a heap.Top-down sift list.elem[start] 
    //to make list[start,...,end] to be a heap with greatest element on the top.
    int child;
    ElemType siftdata;
    siftdata = list.elem[root];  //siftdata is the element need to be adjusted.
    for(child = 2*root+1; child <= end; child = 2*child + 1) {
        //sift down along with the bigger child elements branch.
        if(child < end && list.elem[child] < list.elem[child + 1]) 
            child ++;
        if(siftdata > list.elem[child])
            break;
        list.elem[root] = list.elem[child];
        root = child;  // update root.
    }
    list.elem[root] = siftdata; //insert siftdata at the final position.
}

void HeapSort(UnsortedType &list) { // heapsort
    //Consider list to be a complete binary tree.
    int i;
    ElemType temp;
    //Build heap with greatest element on the top
    for(i = (list.length - 1)/2; i >= 0; i--) {
        SiftDown(list,i,list.length - 1);
    }
    //Pop elements,largest first, into end of list.
    for(i = list.length - 1; i >= 0; i--) {
        // swap the top element with the last element on the heap.
        temp = list.elem[i];
        list.elem[i] = list.elem[0];
        list.elem[0] = temp;
        //siftdown new heap(list[0,i-1])
        SiftDown(list,0,i-1);
    }
}
