package tp

type Triangulation struct {
	nodes []*Point
	// flipper       *Flipper
	// searcher      *Seacher
	// trianglesList *TriangleList
}

const (
	AMOUNT_CLEANING_FACTOR_TRIANGLE_STRUCTURE   float64 = 2.4
	RATIO_DELETING_CONVEX_POINT_FROM_POINT_LIST float64 = 0.2
	MINIMAL_POINTS_FOR_CLEANING                 uint64  = 10000
	AMOUNT_SEARCHER_FACTOR                      float64 = 0.5
)

func (triangulation *Triangulation) getNode(index uint64) *Point {
	return triangulation.nodes[index]
}

func createTriangulation(points []*Point) bool {
	return false
}
