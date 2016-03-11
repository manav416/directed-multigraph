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

// An Edge represents an edge in the directed multi graph.

type Edge struct {
	source   *Vertex
	jobid    int
	dest     *Vertex
	edgeName string
}

func (e Edge) String() string {
	return fmt.Sprintf("%s --> %s (%s)", e.source, e.dest, e.edgeName)
}

func AddEdge(source *Vertex, jobID int, destination *Vertex) *Edge {
	e := &Edge{source, jobID, destination, ""}
	e.edgeName = fmt.Sprintf("%d--%s--%s", jobID, source.name, destination.name)
	source.Out = append(source.Out, e)
	destination.In = append(destination.In, e)

	return e
}

func ShowEdgesEntering(vertex *Vertex) {
	fmt.Println("Finding all edges entering ", vertex)
	for _, e := range vertex.In {
		fmt.Println(e)
	}
}

func ShowEdgesLeaving(vertex *Vertex) {
	fmt.Println("Finding all edges leaving ", vertex)
	for _, e := range vertex.Out {
		fmt.Println(e)
	}
}

func RemoveEdge(edge *Edge) {
	/* First delete this edge from the source */
	removeOutEdge(edge)
	/* Now delete this edge from the destination */
	removeInEdge(edge)
}

// removeOutEdge removes an edge from leaving a vertex
func removeOutEdge(edge *Edge) {
	source := edge.source
	n := len(source.Out)
	for i, e := range source.Out {
		if e == edge {
			// Replace it with the final element and shrink the slice.
			source.Out[i] = source.Out[n-1]
			source.Out[n-1] = nil // Mark this so that the GC can clean it up
			source.Out = source.Out[:n-1]
			return
		}
	}
	panic("edge not found on source: " + edge.String())
}

// removeInEdge removes an edge from entering a vertex
func removeInEdge(edge *Edge) {
	dest := edge.dest
	n := len(dest.In)
	for i, e := range dest.In {
		if e == edge {
			// Replace it with the final element and shrink the slice.
			dest.In[i] = dest.In[n-1]
			dest.In[n-1] = nil // Mark this explicitly so that the GC can pick this up
			dest.In = dest.In[:n-1]
			return
		}
	}
	panic("edge not found on destination: " + edge.String())
}
