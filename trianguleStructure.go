package tp

// Structure of triangles
type TriangleStructure struct {
	IndexNodes, Ribs []uint64
	Triangles        []*TriangleStructure
}
