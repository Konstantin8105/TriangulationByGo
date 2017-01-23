package main

// Structure of triangles
type TriangleStructure struct {
	Nodes, Ribs []int
	Triangles   []*TriangleStructure
}
