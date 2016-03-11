package mgraph

import (
	"fmt"
)

type Graph struct {
	Vertices map[string]*Vertex //Map of vertices by their names
}

// New returns a new Graph
func NewGraph() *Graph {
	g := &Graph{Vertices: make(map[string]*Vertex)}
	return g
}

// CreateVertex returns a Vertex in the graph, creating it if not present.
func (g *Graph) CreateVertex(Nodename string) *Vertex {
	n, ok := g.Vertices[Nodename]
	if !ok {
		n = &Vertex{name: Nodename}
		g.Vertices[Nodename] = n
	}
	return n
}

// A Vertex represents a node in a directed multi graph.
type Vertex struct {
	name string
	In   []*Edge // unordered set of incoming edges
	Out  []*Edge // unordered set of outgoing edges
}

func (n *Vertex) String() string {
	return fmt.Sprintf("vtx:%s", n.name)
}

func (g *Graph) DeleteVertex(n *Vertex) {
	n.deleteIns()
	n.deleteOuts()
	delete(g.Vertices, n.name)
}

// deleteIns deletes all incoming edges to n.
func (n *Vertex) deleteIns() {
	for _, e := range n.In {
		fmt.Println("Removing ", e)
		removeOutEdge(e)
	}
	n.In = nil
}

// deleteOuts deletes all outgoing edges from n.
func (n *Vertex) deleteOuts() {
	for _, e := range n.Out {
		fmt.Println("Removing ", e)
		removeInEdge(e)
	}
	n.Out = nil
}
