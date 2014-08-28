#include "MGraph.h"
#include "stdio.h"
typedef bool PathMatrix[MAX_VERTEX_NUM][MAX_VERTEX_NUM];
typedef EdgeType ShortPathTable[MAX_VERTEX_NUM];

void Shortest_path_DIJ(MGraph G, int v0, PathMatrix &p, ShortPathTable &D){
	//p[v][w]为true，则w是从v0到v的最短路径上的顶点。 (v0 ->...w -> ...-> v)
	//final[v]为true，当且仅当v∈S，即已经求得的v0到v的最短路径
    int i=0,j,v,w,min;
    bool final[MAX_VERTEX_NUM];
    //初始化
    for (v = 0; v < G.vexnum; ++v) {
		final[v] = false;
		D[v] = G.adjmatrix[v0][v];
		for (w = 0; w < G.vexnum; ++w) {
			p[v][w] = false;
		}
		if(D[v] < INFINITY){
			p[v][v0] = true;
			p[v][v] = true;
		}
	}
    D[v0] = 0;  //从v0开始，初始化v0属于S
    final[v0] = true;
    //求出v0到某个v顶点的最短路径，并加v到S中
    for (i = 1; i < G.vexnum; ++i) { //其余n-1个顶点
		min = INFINITY;  //当前所知离v0最近的顶点距离
		for(w = 0;w < G.vexnum;++w){
			if(!final[w]){// w顶点在V-S中
				if(D[w] < min){ //w顶点离v0更近
					v = w;
					min = D[w];
				}
			}
		}
		final[v] = true; //将离v0更近的v加入到S集
		//更新当前最短路径以及距离
		for(w = 0; w < G.vexnum;++w){
			if(!final[w] && G.adjmatrix[v][w] < INFINITY && (min + G.adjmatrix[v][w] < D[w])){
				//修改D[w]以及p[w][j]
				D[w] = min + G.adjmatrix[v][w];
			    for(j = 0;j < G.vexnum;++j){
			    	p[w][j] = p[v][j];
			    	p[w][w] = true;
			    }
			}
		}
	}
}

void DFS(MGraph &G,int v,int i,PathMatrix p) {
    int w;
    //set visited
    G.visit[v] = 1;
    printf("%c ",G.vex[v]);
    for(w = FirstAdjVex(G,v); w >= 0; w = NextAdjVex(G,v,w))
        if(G.visit[w] == 0 && p[i][w] == true)
            DFS(G,w,i,p);
}


void test_DIJ(){
	freopen("input/Graph.txt", "r", stdin);
	MGraph G;
	CreateGraph(G);
	PathMatrix p;
	ShortPathTable D;
	int v0 = 2;
	Shortest_path_DIJ(G,v0,p,D);
	for(int i = 0;i < G.vexnum;++i){
		if(i != v0){
			DFS(G,v0,i,p);
			clearGraph(G);
			printf(" 最短路径长度： %d\n",D[i]);
		}
	}
}
