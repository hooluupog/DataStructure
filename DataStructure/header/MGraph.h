#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE
typedef int Status;
#include "limits.h"

#ifndef _MGRAPH_H_
#define _MGRAPH_H_
#define INFINITY INT_MAX // max value:∞  
#define MAX_VERTEX_NUM 20 // vertex num

typedef enum {DG,DN,UDG,UDN} GraphKind; // { DiGraph,DiNetwork,UndiGraph,UndiNetwork }
typedef char VertexType; // vertices type
typedef int EdgeType; // Unweighted graph or Weighted graph
typedef struct {
    VertexType vex[MAX_VERTEX_NUM]; // list of vertices
    EdgeType adjmatrix[MAX_VERTEX_NUM][MAX_VERTEX_NUM]; // adjacent matrix.
    int visit[MAX_VERTEX_NUM]; // visit mark
   // For unweighted graph:
   // adjmatrix[i][j] = 0, there is no edge between vex[i] and vex[j]. 
   // adjmatrix[i][j] = 1, there is an edge between vex[i] and vex[j]. 
   // For weighted graph:
   // adjmatrix[i][j] = w.(vertex weight.Initial value is INFINITY)
    int vexnum,arcnum; 
    GraphKind kind;
}MGraph;

Status CreateGraph(MGraph &G);
int FirstAdjVex(MGraph G, int v); // find first adjacent vertex of vex[v] and return its index.
int NextAdjVex(MGraph G, int v, int u); // find next adjacent vertex of vex[v] and return its index. u is the beginning position to searching.  
bool isEdge(MGraph G,int v, int u); // edge in (vex[v],vex[u]) or not
int LocateVex(MGraph G,VertexType v); // find index of vertex v in vex[].  
Status BFSTraverse(MGraph G); // breadth-first traverse based on queue
void DFS(MGraph &G,int v); // depth-first search
Status DFSTraverse(MGraph G); // depth-first traverse
void visit(MGraph G,int v); // visit graphic node

#endif
