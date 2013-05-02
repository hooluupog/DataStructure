#include "MGraph.h"
#include "stdio.h"

void testMGraph() {
	/* 
     * direct network adjacent matrix as follows:
     *
	 *  | A B C D E F
     * --------------
     * A| ∞5 ∞7 ∞∞ 
     * B| ∞∞4 ∞∞∞
     * C| 8 ∞∞∞∞9
     * D| ∞∞5 ∞∞6
     * E| ∞∞∞5 ∞∞
     * F| 3 ∞∞∞1 ∞
     *
	 * */
	freopen("input/Graph.txt", "r", stdin);
	printf("MGraph test: \n");
	MGraph G;
	CreateGraph(G);
	printf("Depth first traverse: \n");
	DFSTraverse(G);
	printf("Breadth first traverse: \n");
	BFSTraverse(G);
}
