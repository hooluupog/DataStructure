#include "MGraph.h"
#include "stdio.h"
typedef bool PathMatrix[MAX_VERTEX_NUM][MAX_VERTEX_NUM];
typedef EdgeType DistanceMatrix[MAX_VERTEX_NUM][MAX_VERTEX_NUM];

void ShortestPath_FLOYD(MGraph G, PathMatrix p[], DistanceMatrix &D) {
	//若p[v][w][u]为true，则u是从v到w的当前求得最短路径上的顶点。
	int i, v, u, w;
	for (v = 0; v < G.vexnum; ++v) {
		for (w = 0; w < G.vexnum; ++w) {
			D[v][w] = G.adjmatrix[v][w];
			for (u = 0; u < G.vexnum; ++u) {
				p[v][w][u] = false;
			}
			if (D[v][w] < INFINITY) {
				p[v][w][v] = true;
				p[v][w][w] = true;
			}
		}
	}
	for (u = 0; u < G.vexnum; ++u) {
		for (v = 0; v < G.vexnum; ++v) {
			for (w = 0; w < G.vexnum; ++w) {
				if (D[v][u] < INFINITY && D[u][w] < INFINITY
						&& D[v][u] + D[u][w] < D[v][w]) { //从v经u到w更近
					D[v][w] = D[v][u] + D[u][w];
					for (i = 0; i < G.vexnum; ++i) {
						p[v][w][i] = p[v][u][i] + p[u][w][i];
					}
				}
			}
		}
	}
}

void DFS(MGraph &G, int v0,int v, int i, PathMatrix p[]) {
	// v0到i的路径，v0始终不变， p[v0][i][w] == true表示为最短路径上的结点，递归执行时会进入该结点的分支，v在递归时不断变化
	// 表示每次新的递归开始时新的递归起始结点，w到i的路径。
	int w;
	//set visited
	G.visit[v] = 1;
	printf("%c ", G.vex[v]);
	for (w = FirstAdjVex(G, v); w >= 0; w = NextAdjVex(G, v, w))
		if (G.visit[w] == 0 && p[v0][i][w] == true)
			DFS(G,v0, w, i, p);
}

void test_FLOYD() {
	freopen("input/Graph.txt", "r", stdin);
	MGraph G;
	CreateGraph(G);
	PathMatrix p[G.vexnum];
	DistanceMatrix D;
	ShortestPath_FLOYD(G, p, D);
	for (int v = 0; v < G.vexnum; ++v) {
		for (int i = 0; i < G.vexnum; ++i) {
			if (i != v) {
				DFS(G, v, v, i, p);
				clearGraph(G);
				printf(" 最短路径长度： %d\n", D[v][i]);
			}
		}
	}
}
