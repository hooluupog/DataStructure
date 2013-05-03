/**
** 让任务管理器中的cpu曲线呈现出正弦曲线
**/
#include "windows.h"
#include "stdlib.h"
#include "math.h"

#define SPLIT 0.01
#define COUNT 200
#define INTERVAL 300
#define PI 3.141592625

int main()
{
	DWORD busySpan[COUNT];
	DWORD idleSpan[COUNT];
	int half = INTERVAL/2;
	double radian = 0.0;
	DWORD startTime = 0;
	int j = 0;
	int i;    
	for(i = 0;i < COUNT;++i)
	{
		busySpan[i] = half + sin(PI*radian)*half;
		idleSpan[i] = INTERVAL - busySpan[i];
		radian += SPLIT;
	}
	while(TRUE)
	{
		j = j%COUNT;
		startTime = GetTickCount();
		while(GetTickCount() - startTime <= busySpan[j]);
		Sleep(idleSpan[j]);
		j++;
	}
	return 0;
}
