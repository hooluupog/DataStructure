#define OK 1
#define ERROR 0
#define _CRT_SECURE_NO_DEPRECATE
typedef int Status;

#ifndef _ALGRAPH_H_
#define _ALGRAPH_H_
#define MAX_VERTEX_NUM 20 // vertex num

typedef enum {DG,DN,UDG,UDN} GraphKind; // { DiGraph,DiNetwork,UndiGraph,UndiNetwork }
typedef char VertexType; // vertices data type
typedef int EdgeType; // Unweighted graph or Weighted graph

typedef struct ArcNode{ //tail node of the arc. e.g. arc ~ (x,y), y is tail and x is head of the arc(edge).
    int adjvex; //the index of arc node
    struct ArcNode *nextarc; // point to next arc node
    EdgeType weight; // for graph,weight=0; for network,it represents edge's weight.
}ArcNode;

typedef struct VNode{ // vertices
    ArcNode *firstarc; // point to the first adjacent node in its adjacent list.
    VertexType data; // vertice data
    int visited;  //visit mark
}VNode,AdjList[MAX_VERTEX_NUM];

typedef struct{
    AdjList vertices; 
    int vexnum,arcnum; 
    GraphKind kind;
}ALGraph;

Status CreateGraph(ALGraph &G);
int FirstAdjVex(ALGraph G, int v); // find first adjacent vertex of vex[v] and return its index.
int NextAdjVex(ALGraph G, int v, int u); // find next adjacent vertex of vex[v] and return its index. u is the beginning position to searching.  
bool isEdge(ALGraph G,int v, int u); // edge in (vex[v],vex[u]) or not
int LocateVex(ALGraph G,VertexType v); // find index of vertex v in vex[].  
void DFS(ALGraph &G,int v); // depth-first search
Status DFSTraverse(ALGraph G); // depth-first traverse
Status BFSTraverse(ALGraph G); // breadth-first traverse based on queue
void visit(ALGraph G,int v); // visit graphic node

#endif
