#include "SqList.h"
typedef SqList UnsortedType;
#ifndef _ALGORITHM_H
#define _ALGORITHM_H
void Qsort(UnsortedType &list,int low,int high); // quicksort
void HeapSort(UnsortedType &list); // heapsort
#endif
