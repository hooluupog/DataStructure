#include "ALGraph.h"
#include "stdio.h"
#include "SqStack.h"
Status TopoSort(ALGraph G){//拓扑排序
	SqStack S;
	int count,k,i;
	ArcNode *p;
	char indegree[MAX_VERTEX_NUM];
	FindInDegree(G,indegree); //对各结点求入度
	InitStack_Sq(S);
	for(i = 0;i < G.vexnum;++i){
		if(!indegree[i]){
			Push_Sq(S,i);//入度为0者进栈
		}
	}
	count = 0;
	while(!IsEmpty_Sq(S)){
		Pop_Sq(S,i);
		printf("%c",G.vertices[i].data);
		count++; //输出i号顶点并计数
		for(p = G.vertices[i].firstarc;p != NULL;p = p->nextarc){
			k = p->adjvex; //对i号顶点的每个邻接顶点的入度减1
			if(!(--indegree[k])){
				Push_Sq(S,k); //若入度减为0，则入栈
			}
		}
	}
	if(count < G.vexnum) return ERROR;//存在回路
	else return OK;
}

void TopoSort_DFS(ALGraph &G,int v,SqStack &S){
	int w;
	//set visited
	G.vertices[v].visited = 1; // 标记已访问避免重复入栈
	for(w = FirstAdjVex(G,v); w >= 0; w = NextAdjVex(G,v,w)){
		if(G.vertices[w].visited == 0){
			TopoSort_DFS(G,w,S);
		}
	}
	Push_Sq(S,v);
}

Status Topo_DFSTraverse(ALGraph G,SqStack &S){
	int i;
	for(i = 0; i < G.vexnum;i++)
		if(G.vertices[i].visited == 0)
			TopoSort_DFS(G,i,S);
	return OK;
}

void test_TopoSort() {
	freopen("input/TopoLogicalSort2.txt", "r", stdin);
	ALGraph G;
	SqStack S;
	int v;
	CreateGraph(G);
    TopoSort(G);
    printf("\n");
    InitStack_Sq(S);
    Topo_DFSTraverse(G,S);
    while(!(IsEmpty_Sq(S))){
    	Pop_Sq(S,v);
    	printf("%c",G.vertices[v].data);
    }
}
