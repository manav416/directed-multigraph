package main

import (
	"directed-multigraph/mgraph"
)

type VertexInfo struct {
	vertexName string
	vertexCost uint64
}

type EdgeInfo struct {
	edgeName string
	edgeData interface{}
}

func main() {
	/* Create a graph in which all the vertices
	 * and edges will be built */

	var dummyVertexInfo VertexInfo
	var dummyEdgeData EdgeInfo

	newGraph := mgraph.NewGraph()
	n1 := newGraph.CreateVertex("hop1", dummyVertexInfo)
	n2 := newGraph.CreateVertex("hop2", dummyVertexInfo)
	n3 := newGraph.CreateVertex("source-1", dummyVertexInfo)
	n4 := newGraph.CreateVertex("source-2", dummyVertexInfo)
	n5 := newGraph.CreateVertex("destination-1", dummyVertexInfo)
	n6 := newGraph.CreateVertex("destination-2", "Some String")

	/* Two ways to create edges */
	e1 := mgraph.AddEdge(n3, 666, n1)
	mgraph.AddEdge(n4, "Some Other Data", n1)
	mgraph.AddEdge(n1, dummyEdgeData, n2)
	mgraph.AddEdge(n2, dummyEdgeData, n5)
	mgraph.AddEdge(n2, dummyEdgeData, n6)

	mgraph.ShowEdgesEntering(n1)
	mgraph.RemoveEdge(e1)
	mgraph.ShowEdgesEntering(n1)

	newGraph.DeleteVertex(n1)

	mgraph.ShowEdgesEntering(n1)

}
