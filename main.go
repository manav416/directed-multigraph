package main

import (
	"directed-multigraph/mgraph"
)

func main() {
	/* Create a graph in which all the vertices
	 * and edges will be built */

	newGraph := mgraph.NewGraph()
	n1 := newGraph.CreateVertex("hop1")
	n2 := newGraph.CreateVertex("hop2")
	n3 := newGraph.CreateVertex("source-1")
	n4 := newGraph.CreateVertex("source-2")
	n5 := newGraph.CreateVertex("destination-1")
	n6 := newGraph.CreateVertex("destination-2")

	/* Two ways to create edges */
	e1 := mgraph.AddEdge(n3, 666, n1)
	mgraph.AddEdge(n4, 667, n1)
	mgraph.AddEdge(n1, 666, n2)
	mgraph.AddEdge(n2, 666, n5)
	mgraph.AddEdge(n2, 667, n6)

	mgraph.ShowEdgesEntering(n1)
	mgraph.RemoveEdge(e1)
	mgraph.ShowEdgesEntering(n1)

	newGraph.DeleteVertex(n1)

	mgraph.ShowEdgesEntering(n1)

}
