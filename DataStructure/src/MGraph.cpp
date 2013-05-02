#include "MGraph.h"
#include "fatal.h"
#include "stdio.h"
#include "LinkQueue.h"

Status CreateGraph(MGraph &G){
    int i,j,k;
    VertexType u,v;
    EdgeType weight;
    //G.kind---- 0:DG 1:DN 2:UDG 3:UDN
    scanf("%d%d%d\n",&G.vexnum,&G.arcnum,&G.kind);
    if(G.kind > 3 || G.kind < 0){
        Error(" Incorrect graph kind input. the legal number must between 0 and 3.");
        return ERROR;
    }
    for(i = 0; i < G.vexnum; i++) {// Construct vertex nodes and init adjacent matrix.
        scanf("%c\n", &G.vex[i]);
        G.visit[i] = 0; //set unvisited
        for(j = 0; j < G.vexnum; j++) {
            if(G.kind == DN || G.kind == UDN)
                G.adjmatrix[i][j] = INFINITY; 
            else
                G.adjmatrix[i][j] = 0; 
        }//for 
    }//for
    for(k = 0; k < G.arcnum; k++) { //update arc info in adjacent matrix.
        scanf("%c %c\n",&u,&v);
        //locate index of u and v
        i = LocateVex(G,u);
        j = LocateVex(G,v);
        if(G.kind == DG || G.kind == UDG)
            G.adjmatrix[i][j] = 1; 
        else{ // network,need input weight info.
            scanf("%d\n",&weight);
            G.adjmatrix[i][j] = weight;
        }//else
        if(G.kind == UDG || G.kind == UDN)
            G.adjmatrix[j][i] = G.adjmatrix[i][j]; // the matrix is symmetry in undirect graph or undirect network.
    }//for
    return OK;
}//CreateGraph 

// find first adjacent vertex of vex[v] and return its index.
int FirstAdjVex(MGraph G, int v) {
    int j;
    for(j = 0; j < G.vexnum;j++)
        if(G.adjmatrix[v][j] != 0 && G.adjmatrix[v][j] != INFINITY)
            return j;
   return -1; 
} 
    
// find next adjacent vertex of vex[v] and return its index. u is the beginning position to searching.  
int NextAdjVex(MGraph G, int v, int u) {
    int j;
    for(j = u+1; j < G.vexnum;j++)
        if(G.adjmatrix[v][j] != 0 && G.adjmatrix[v][j] != INFINITY)
            return j;
   return -1; 
} 

// edge in (vex[v],vex[u]) or not
bool isEdge(MGraph G,int v, int u) {
    int w;
    w = FirstAdjVex(G,v);
    for(; w >= 0; w = NextAdjVex(G,v,w))
        if(w == u)
            return true;
    return false;
} 

// find index of vertex v in vex[].  
int LocateVex(MGraph G,VertexType v){
    int i;
    for(i = 0; i < G.vexnum; i++)
        if(G.vex[i] == v)
            return i;
    return -1;
} 

// depth-first search
void DFS(MGraph &G,int v) {
    int w;
    //set visited
    G.visit[v] = 1;
    visit(G,v);
    for(w = FirstAdjVex(G,v); w >= 0; w = NextAdjVex(G,v,w))
        if(G.visit[w] == 0)
            DFS(G,w);
} 

// depth-first traverse
Status DFSTraverse(MGraph G){
	int i;
	for(i = 0; i < G.vexnum;i++)
		if(G.visit[i] == 0)
			DFS(G,i);
	return OK;
}

// breadth-first traverse based on queue
Status BFSTraverse(MGraph G) {
    LinkQueue Q;
    Init_LinkQueue(Q);
    int w,i,u;
    for(i = 0; i < G.vexnum; i++) //vertex nodes
        if(G.visit[i] == 0){ // unvisited
            G.visit[i] = 1; // set visited
            visit(G,i); // visit it
            En_LinkQueue(Q,i); // vex[i] enqueue
            while(!IsEmpty_LinkQueue(Q)) {
                De_LinkQueue(Q,u);
                for(w = FirstAdjVex(G,w); w >= 0; w = NextAdjVex(G,u,w))
                    if(G.visit[w] == 0){ // visit adjacent nodes
                        G.visit[w] = 1; 
                        visit(G,w);
                        En_LinkQueue(Q,w);
                    }//if
            }//while
        }//if
    Destroy_LinkQueue(Q); // free memory space
    return OK;
} 

void visit(MGraph G,int v) {
    printf("%c\n",G.vex[v]);
}
