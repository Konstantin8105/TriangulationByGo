package tp

import (
	"container/list"
	"fmt"

	"github.com/Konstantin8105/tp/bb"
	"github.com/Konstantin8105/tp/point"
)

// main data structure for triangulation
//
// Triangulation data structure  "Nodes, simple ribs и triangles"
// book "Algoritm building and analyse triangulation", A.B.Skvorcov.
// paragraph 1.2.5
//
type data struct {
	nodes     [3]int   // indexes of triangle points
	triangles [3]*data // indexes of near triangles
	ribs      [3]int   // indexes of triangle ribs
}

func (d *data) changeClockwise() {
	d.nodes[0], d.nodes[1] = d.nodes[1], d.nodes[0]
	d.ribs[1], d.ribs[2] = d.ribs[2], d.ribs[1]
	d.triangles[1], d.triangles[2] = d.triangles[2], d.triangles[1]
}

type Triangulation struct {
	ps []point.Point
	ds *list.List

	// last used triangle
	last *data
}

func (tr *Triangulation) New(ps ...point.Point) error {

	// find border box
	b := bb.New()
	for i := range ps {
		b.Add(ps[i])
	}

	//
	// create pseudo-box.
	// all points must be inside pseudo-box
	//
	//	P1     P2
	//	o---1---o
	//	|      /|
	//	|  0  / |
	//	|    /  |
	//	0   2   3
	//	|  /    |
	//	| /  1  |
	//  |/      |
	//  o---4---o
	//	P0     P3
	//
	pps := []point.Point{ // pseudo-box points
		point.Point{X: b.Xmin - 1.0, Y: b.Ymin - 1.0}, // P0
		point.Point{X: b.Xmin - 1.0, Y: b.Ymax + 1.0}, // P1
		point.Point{X: b.Xmax + 1.0, Y: b.Ymax + 1.0}, // P2
		point.Point{X: b.Xmax + 1.0, Y: b.Ymin - 1.0}, // P3
	}
	defer func() {
		for i := range pps {
			tr.remove(pps[i])
		}
	}()
	tr.ps = append(pps, ps...)

	//
	// create points, ribs, triangles pseudo-box
	//
	t0 := data{
		nodes: [3]int{0, 1, 2},
		ribs:  [3]int{0, 1, 2},
	}
	t1 := data{
		nodes: [3]int{2, 3, 0},
		ribs:  [3]int{3, 4, 2},
	}
	t0.triangles[2] = &t1
	t1.triangles[2] = &t0
	tr.ds.PushFront(&t0)
	tr.ds.PushFront(&t1)

	//
	// add points in triangles
	//
	tr.last = &t0
	for i := 5; i < len(tr.ps); i++ {
		if err := tr.add(i); err != nil {
			return err
		}
	}
	return nil
}

func (tr *Triangulation) remove(p point.Point) error {
	panic("remove")
}

func (tr *Triangulation) add(next int) error {
	searcher(next)
	state := movingByConvexHull(next)
	err := fmt.Errorf("Strange point #%d : %s", next, tr.ps[next])
	switch state {
	case pointInside:
		err = addNextPointInTriangle(nextPoint)
	case pointOnLine0:
		err = addNextPointOnLine(nextPoint, 0)
	case pointOnLine1:
		err = addNextPointOnLine(nextPoint, 1)
	case pointOnLine2:
		err = addNextPointOnLine(nextPoint, 2)
	case pointOnCorner:
		err = nil
	}
	return
}

func (tr *Triangulation) movingByConvexHull(Point point) state {
	value := [3]POINT_LINE_STATE{}
	beginTriangle := tr.last
	for {
		//add reserve searching
		value[0] = calculateValuePointOnLine(getNode(beginTriangle.iNodes[0]), getNode(beginTriangle.iNodes[1]), point)
		if Geometry.isAtRightOf(value[0]) {
			beginTriangle = beginTriangle.triangles[0]
		} else {
			whichOp := 0
			value[1] = calculateValuePointOnLine(getNode(beginTriangle.iNodes[1]), getNode(beginTriangle.iNodes[2]), point)
			if Geometry.isAtRightOf(value[1]) {
				whichOp += 1
			}
			value[2] = calculateValuePointOnLine(getNode(beginTriangle.iNodes[2]), getNode(beginTriangle.iNodes[0]), point)
			if Geometry.isAtRightOf(value[2]) {
				whichOp += 2
			}

			switch whichOp {
			case 0:
			case 1:
				beginTriangle = beginTriangle.triangles[1]
			case 2:
				beginTriangle = beginTriangle.triangles[2]
			default:
				if distanceLineAndPoint(getNode(beginTriangle.iNodes[1]), getNode(beginTriangle.iNodes[2]), point) >
					distanceLineAndPoint(getNode(beginTriangle.iNodes[2]), getNode(beginTriangle.iNodes[0]), point) {
					beginTriangle = beginTriangle.triangles[1]
				} else {
					beginTriangle = beginTriangle.triangles[2]
				}
			}
		}
	}
	trianglePoint = []point.Point{
		getNode(beginTriangle.iNodes[0]),
		getNode(beginTriangle.iNodes[1]),
		getNode(beginTriangle.iNodes[2]),
	}
	setSearcher(beginTriangle)
	return statePointInTriangle(point, trianglePoint, value)
}
