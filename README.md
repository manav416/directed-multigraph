# directed-multigraph
Directed Multigraph Package in Golang

In mathematics, and more specifically in graph theory, a multigraph is a graph which is permitted to have multiple edges (also called parallel edges[1]), that is, edges that have the same end nodes. Thus two vertices may be connected by more than one edge.

This is a simple Golang package to use directed multigraph

You can create a graph in which you can add vertices. 
There are APIs to create edges that can connect to multiple vertices.
Each edge has a jobid field which can be replaced by an empty interface
or some structure of your choice. This can be used to define thin or thick edges
connecting the vertices.

There is a sample main.go which illustrates how this package can be used.

Please get back to me with your comments on manavbhatia@gmail.com
