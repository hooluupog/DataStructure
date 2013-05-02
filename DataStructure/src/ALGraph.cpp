#include "ALGraph.h"
#include "fatal.h"
#include "stdio.h"
#include "LinkQueue.h"

Status CreateGraph(ALGraph &G){
    //G.kind---- 0:DG 1:DN 2:UDG 3:UDN
    scanf("%d%d%d\n",&G.vexnum,&G.arcnum,&G.kind);
    if(G.kind > 3 || G.kind < 0){
        Error(" Incorrect graph kind input. the legal number must between 0 and 3.");
        return ERROR;
    }
    int i,j,k;
    EdgeType weight;
    ArcNode *pi,*pj;
    VertexType sv,ev;
    for(i = 0; i < G.vexnum; i++) { //Construct vertex nodes.Input vertex data and init adjlist head pointer.
        scanf("%c\n", &G.vertices[i].data); 
        G.vertices[i].firstarc = NULL;
        G.vertices[i].visited = 0;  //set unvisited
    }//for
    for(k = 0; k < G.arcnum; k++) { //Construct arc.Input head and tail of an edge.
        scanf("%c %c\n",&sv,&ev); 
        // locate index of sv and ev
        i = LocateVex(G,sv);
        j = LocateVex(G,ev);
        pi = (ArcNode *)malloc(sizeof(ArcNode));
        if(!pi){
            Error("out of space");
            return ERROR;
        }//if
        pi->adjvex = j;  
        // edge's weight info
        if(G.kind == DN || G.kind == UDN) {
            scanf("%d\n",&weight);
        }
        else 
            weight = 0;
        //Head insert to construct vertices[i]'s adjlist.
        pi->weight = weight;
        pi->nextarc = G.vertices[i].firstarc;
        G.vertices[i].firstarc = pi;
        if(G.kind == UDG || G.kind == UDN) {
            //Construct vertices[j]'s adjlist in the light of symmetry in UDG/UDN.
            pj = (ArcNode *)malloc(sizeof(ArcNode));
            if(!pj){
            Error("out of space");
            return ERROR;
            }//if
            pj->adjvex = i;
            pj->weight = weight;
            pj->nextarc = G.vertices[j].firstarc;
            G.vertices[j].firstarc = pj;
        }//if
    }//for
    return OK;
}//CreateGraph 

// find first adjacent vertex of vex[v] and return its index.
int FirstAdjVex(ALGraph G, int v) {
    return G.vertices[v].firstarc == NULL ? -1 : G.vertices[v].firstarc->adjvex;
} 

// find next adjacent vertex of vex[v] and return its index. u is the beginning position to searching.  
int NextAdjVex(ALGraph G, int v, int u) {
    ArcNode *p;
    for(p = G.vertices[v].firstarc; p != NULL && p->adjvex != u; p = p->nextarc)
        ;
    if(p) 
        p = p->nextarc;
    return p != NULL ? p->adjvex : -1;
} 

// exist and edge ~ (vex[v],vex[u]) or not
bool isEdge(ALGraph G,int v, int u) {
    int w;
    w = FirstAdjVex(G,v);
    for(; w >= 0; w = NextAdjVex(G,v,w))
        if(w == u)
            return true;
    return false;
} 

// find index of vertex v in vex[].  
int LocateVex(ALGraph G,VertexType v){
    int i;
    for(i = 0; i < G.vexnum; i++)
        if(G.vertices[i].data == v)
            return i;
    return -1;
} 

// depth-first search
void DFS(ALGraph &G,int v) {
    int w;
    //set visited
    G.vertices[v].visited = 1;
    visit(G,v);
    for(w = FirstAdjVex(G,v); w >= 0; w = NextAdjVex(G,v,w))
        if(G.vertices[w].visited == 0)
            DFS(G,w);
} 

// depth-first traverse
Status DFSTraverse(ALGraph G){
	int i;
	for(i = 0; i < G.vexnum;i++)
		if(G.vertices[i].visited == 0)
			DFS(G,i);
	return OK;
}

// breadth-first search based on queue
Status BFSTraverse(ALGraph G) {
    LinkQueue Q;
    Init_LinkQueue(Q);
    int w,i,u;
    for(i = 0; i < G.vexnum; i++) //vertex nodes
        if(G.vertices[i].visited == 0){ // unvisited
            G.vertices[i].visited = 1; // set visited
            visit(G,i); // visit it
            En_LinkQueue(Q,i); // vex[i] enqueue
            while(!IsEmpty_LinkQueue(Q)) {
                De_LinkQueue(Q,u);
                for(w = FirstAdjVex(G,u); w >= 0; w = NextAdjVex(G,u,w))
                    if(G.vertices[w].visited == 0){ // visit adjacent nodes
                        G.vertices[w].visited = 1; 
                        visit(G,w);
                        En_LinkQueue(Q,w);
                    }//if
            }//while
        }//if
    Destroy_LinkQueue(Q); // free memory space
    return OK;
} 

//visit graphic node
void visit(ALGraph G,int v) {
    printf("%c\n",G.vertices[v].data);
}
