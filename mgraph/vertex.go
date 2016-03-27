/*
Copyright 2016 Manav Bhatia

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mgraph

import (
	"fmt"
)

// A Vertex represents a node in a directed multi graph.
type Vertex struct {
	name       string
	In         []*Edge     // unordered set of incoming edges
	Out        []*Edge     // unordered set of outgoing edges
	vertexData interface{} // store information specific to the Vertex
}

type Graph struct {
	Vertices map[string]*Vertex //Map of vertices by their names
}

// New returns a new Graph
func NewGraph() *Graph {
	g := &Graph{Vertices: make(map[string]*Vertex)}
	return g
}

// CreateVertex returns a Vertex in the graph, creating it if not present.
func (g *Graph) CreateVertex(Nodename string, data interface{}) *Vertex {
	n, ok := g.Vertices[Nodename]
	if !ok {
		n = &Vertex{name: Nodename, vertexData: data}
		g.Vertices[Nodename] = n
	}
	return n
}

func (n *Vertex) String() string {
	return fmt.Sprintf("<%s>", n.name)
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

func (n Vertex) GetData() interface{} {
	return n.vertexData
}

func (n Vertex) GetVertexName() string {
	return n.name
}
