#include "ALGraph.h"
#include "stdio.h"

void testALGraph() {
	/* 
     * direct network adjacent list as follows:
     *
	 *  | A B C D E F
     * --------------
     * A| ∞5 ∞7 ∞∞          |A|--->D--->B 
     * B| ∞∞4 ∞∞∞          |B|--->C
     * C| 8 ∞∞∞∞9  ----->   |C|--->F--->A
     * D| ∞∞5 ∞∞6           |D|--->F--->C 
     * E| ∞∞∞5 ∞∞          |E|--->D
     * F| 3 ∞∞∞1 ∞          |F|--->E--->A
     *
	 * */
	freopen("input/Graph.txt","r",stdin);
	printf("ALGraph test: \n");
	ALGraph G;
	CreateGraph(G);
	printf("Depth first traverse: \n");
    DFSTraverse(G);
    printf("Breadth first traverse: \n");
    BFSTraverse(G);
}
