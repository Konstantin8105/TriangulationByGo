package tp

import "github.com/Konstantin8105/tp/point"

type Triangulation struct {
	nodes []*point.Point
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

func (triangulation *Triangulation) getNode(index uint64) *point.Point {
	return triangulation.nodes[index]
}

func Run(points []point.Point) bool {
	return false
}
