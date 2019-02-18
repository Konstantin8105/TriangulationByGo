package tp

// Structure of triangles
type TriangleStructure struct {
	IndexNodes, Ribs []uint64
	Triangles        []*TriangleStructure
}

// func changeClockwise() {
// int temp;
// temp = iNodes[0];
// iNodes[0] = iNodes[1];
// iNodes[1] = temp;
// temp = iRibs[1];
// iRibs[1] = iRibs[2];
// iRibs[2] = temp;
// TriangleStructure tri = triangles[1];
// triangles[1] = triangles[2];
// triangles[2] = tri;
// }
